/*
 * @Date: 2022-04-14 15:27:45
 * @LastEditors: recar
 * @LastEditTime: 2022-04-15 14:51:05
 */

// copy https://github.com/jweny/pocassist/blob/master/pkg/cel/cel.go
package cel

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
	"woodpecker/pkg/cel/proto"
	reverse2 "woodpecker/pkg/cel/reverse"
	"woodpecker/pkg/util"

	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter/functions"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

//	判断s1是否包含s2
var containsFunc = &functions.Overload{
	Operator: "contains_string",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		v1, ok := lhs.(types.String)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to contains", lhs.Type())
		}
		v2, ok := rhs.(types.String)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to contains", rhs.Type())
		}
		return types.Bool(strings.Contains(string(v1), string(v2)))
	},
}

// 判断s1是否包含s2, 忽略大小写
var iContainsDec = decls.NewFunction("icontains", decls.NewInstanceOverload("string_icontains_string", []*exprpb.Type{decls.String, decls.String}, decls.Bool))
var iContainsFunc = &functions.Overload{
	Operator: "string_icontains_string",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		v1, ok := lhs.(types.String)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to icontains", lhs.Type())
		}
		v2, ok := rhs.(types.String)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to icontains", rhs.Type())
		}
		return types.Bool(strings.Contains(strings.ToLower(string(v1)), strings.ToLower(string(v2))))
	},
}

//	判断b1 是否包含 b2
var bcontainsDec = decls.NewFunction("bcontains", decls.NewInstanceOverload("bytes_bcontains_bytes", []*exprpb.Type{decls.Bytes, decls.Bytes}, decls.Bool))
var bcontainsFunc = &functions.Overload{
	Operator: "bytes_bcontains_bytes",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		v1, ok := lhs.(types.Bytes)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to bcontains", lhs.Type())
		}
		v2, ok := rhs.(types.Bytes)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to bcontains", rhs.Type())
		}
		return types.Bool(bytes.Contains(v1, v2))
	},
}

//	使用正则表达式s1来匹配s2
var matchFunc = &functions.Overload{
	Operator: "matches_string",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		v1, ok := lhs.(types.String)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to match", lhs.Type())
		}
		v2, ok := rhs.(types.String)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to match", rhs.Type())
		}
		ok, err := regexp.Match(string(v1), []byte(v2))
		if err != nil {
			return types.NewErr("%v", err)
		}
		return types.Bool(ok)
	},
}

//	使用正则表达式s1 来 匹配b1
var bmatchDec = decls.NewFunction("bmatches",
	decls.NewInstanceOverload("string_bmatch_bytes",
		[]*exprpb.Type{decls.String, decls.Bytes},
		decls.Bool))

var bmatchFunc = &functions.Overload{
	Operator: "string_bmatch_bytes",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		v1, ok := lhs.(types.String)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to bmatch", lhs.Type())
		}
		v2, ok := rhs.(types.Bytes)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to bmatch", rhs.Type())
		}
		ok, err := regexp.Match(string(v1), v2)
		if err != nil {
			return types.NewErr("%v", err)
		}
		return types.Bool(ok)
	},
}

//	map 中是否包含某个 key，目前只有 headers 是 map[string][string] 类型
var inDec = decls.NewFunction("in", decls.NewInstanceOverload("string_in_map_key", []*exprpb.Type{decls.String, decls.NewMapType(decls.String, decls.String)}, decls.Bool))
var inFunc = &functions.Overload{
	Operator: "string_in_map_key",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		v1, ok := lhs.(types.String)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to in", lhs.Type())
		}
		v2, ok := rhs.(types.Bytes)
		// 临时方案 判断字符串是否在 map中
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to in", lhs.Type())
		}
		ok = strings.Contains(string(v2), string(v1))
		return types.Bool(ok)
	},
}

//  字符串的 md5
var md5Dec = decls.NewFunction("md5", decls.NewOverload("md5_string", []*exprpb.Type{decls.String}, decls.String))
var md5Func = &functions.Overload{
	Operator: "md5_string",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.String)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to md5_string", value.Type())
		}
		return types.String(fmt.Sprintf("%x", md5.Sum([]byte(v))))
	},
}

//	两个范围内的随机数
var randomIntDec = decls.NewFunction("randomInt", decls.NewOverload("randomInt_int_int", []*exprpb.Type{decls.Int, decls.Int}, decls.Int))
var randomIntFunc = &functions.Overload{
	Operator: "randomInt_int_int",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		from, ok := lhs.(types.Int)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to randomInt", lhs.Type())
		}
		to, ok := rhs.(types.Int)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to randomInt", rhs.Type())
		}
		min, max := int(from), int(to)
		return types.Int(rand.Intn(max-min) + min)
	},
}

//	指定长度的小写字母组成的随机字符串
var randomLowercaseDec = decls.NewFunction("randomLowercase", decls.NewOverload("randomLowercase_int", []*exprpb.Type{decls.Int}, decls.String))
var randomLowercaseFunc = &functions.Overload{
	Operator: "randomLowercase_int",
	Unary: func(value ref.Val) ref.Val {
		n, ok := value.(types.Int)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to randomLowercase", value.Type())
		}
		return types.String(util.RandLetters(int(n)))
	},
}

//	将字符串进行 base64 编码
var base64StringDec = decls.NewFunction("base64", decls.NewOverload("base64_string", []*exprpb.Type{decls.String}, decls.String))
var base64StringFunc = &functions.Overload{
	Operator: "base64_string",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.String)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to base64_string", value.Type())
		}
		return types.String(base64.StdEncoding.EncodeToString([]byte(v)))
	},
}

//	将bytes进行 base64 编码
var base64BytesDec = decls.NewFunction("base64", decls.NewOverload("base64_bytes", []*exprpb.Type{decls.Bytes}, decls.String))
var base64BytesFunc = &functions.Overload{
	Operator: "base64_bytes",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.Bytes)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to base64_bytes", value.Type())
		}
		return types.String(base64.StdEncoding.EncodeToString(v))
	},
}

//	将字符串进行 base64 解码
var base64DecodeStringDec = decls.NewFunction("base64Decode", decls.NewOverload("base64Decode_string", []*exprpb.Type{decls.String}, decls.String))
var base64DecodeStringFunc = &functions.Overload{
	Operator: "base64Decode_string",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.String)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to base64Decode_string", value.Type())
		}
		decodeBytes, err := base64.StdEncoding.DecodeString(string(v))
		if err != nil {
			return types.NewErr("%v", err)
		}
		return types.String(decodeBytes)
	},
}

//	将bytes进行 base64 编码
var base64DecodeBytesDec = decls.NewFunction("base64Decode", decls.NewOverload("base64Decode_bytes", []*exprpb.Type{decls.Bytes}, decls.String))
var base64DecodeBytesFunc = &functions.Overload{
	Operator: "base64Decode_bytes",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.Bytes)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to base64Decode_bytes", value.Type())
		}
		decodeBytes, err := base64.StdEncoding.DecodeString(string(v))
		if err != nil {
			return types.NewErr("%v", err)
		}
		return types.String(decodeBytes)
	},
}

//	将字符串进行 urlencode 编码
var urlencodeStringDec = decls.NewFunction("urlencode", decls.NewOverload("urlencode_string", []*exprpb.Type{decls.String}, decls.String))
var urlencodeStringFunc = &functions.Overload{
	Operator: "urlencode_string",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.String)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to urlencode_string", value.Type())
		}
		return types.String(url.QueryEscape(string(v)))
	},
}

//	将bytes进行 urlencode 编码
var urlencodeBytesDec = decls.NewFunction("urlencode", decls.NewOverload("urlencode_bytes", []*exprpb.Type{decls.Bytes}, decls.String))
var urlencodeBytesFunc = &functions.Overload{
	Operator: "urlencode_bytes",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.Bytes)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to urlencode_bytes", value.Type())
		}
		return types.String(url.QueryEscape(string(v)))
	},
}

//	将字符串进行 urldecode 解码
var urldecodeStringDec = decls.NewFunction("urldecode", decls.NewOverload("urldecode_string", []*exprpb.Type{decls.String}, decls.String))
var urldecodeStringFunc = &functions.Overload{
	Operator: "urldecode_string",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.String)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to urldecode_string", value.Type())
		}
		decodeString, err := url.QueryUnescape(string(v))
		if err != nil {
			return types.NewErr("%v", err)
		}
		return types.String(decodeString)
	},
}

//	将 bytes 进行 urldecode 解码
var urldecodeBytesDec = decls.NewFunction("urldecode", decls.NewOverload("urldecode_bytes", []*exprpb.Type{decls.Bytes}, decls.String))
var urldecodeBytesFunc = &functions.Overload{
	Operator: "urldecode_bytes",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.Bytes)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to urldecode_bytes", value.Type())
		}
		decodeString, err := url.QueryUnescape(string(v))
		if err != nil {
			return types.NewErr("%v", err)
		}
		return types.String(decodeString)
	},
}

//	截取字符串
var substrDec = decls.NewFunction("substr", decls.NewOverload("substr_string_int_int", []*exprpb.Type{decls.String, decls.Int, decls.Int}, decls.String))
var substrFunc = &functions.Overload{
	Operator: "substr_string_int_int",
	Function: func(values ...ref.Val) ref.Val {
		if len(values) == 3 {
			str, ok := values[0].(types.String)
			if !ok {
				return types.NewErr("invalid string to 'substr'")
			}
			start, ok := values[1].(types.Int)
			if !ok {
				return types.NewErr("invalid start to 'substr'")
			}
			length, ok := values[2].(types.Int)
			if !ok {
				return types.NewErr("invalid length to 'substr'")
			}
			runes := []rune(str)
			if start < 0 || length < 0 || int(start+length) > len(runes) {
				return types.NewErr("invalid start or length to 'substr'")
			}
			return types.String(runes[start : start+length])
		} else {
			return types.NewErr("too many arguments to 'substr'")
		}
	},
}

//	暂停执行等待指定的秒数
var sleepDec = decls.NewFunction("sleep", decls.NewOverload("sleep_int", []*exprpb.Type{decls.Int}, decls.Bool))
var sleepFunc = &functions.Overload{
	Operator: "sleep_int",
	Unary: func(value ref.Val) ref.Val {
		v, ok := value.(types.Int)
		if !ok {
			return types.ValOrErr(value, "unexpected type '%v' passed to sleep", value.Type())
		}
		time.Sleep(time.Duration(v) * time.Second)
		return types.Bool(true)
	},
}

//	反连平台结果
var reverseWaitDec = decls.NewFunction("wait", decls.NewInstanceOverload("reverse_wait_int", []*exprpb.Type{decls.Any, decls.Int}, decls.Bool))
var reverseWaitFunc = &functions.Overload{
	Operator: "reverse_wait_int",
	Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
		reverse, ok := lhs.Value().(*proto.Reverse)
		if !ok {
			return types.ValOrErr(lhs, "unexpected type '%v' passed to 'wait'", lhs.Type())
		}
		timeout, ok := rhs.Value().(int64)
		if !ok {
			return types.ValOrErr(rhs, "unexpected type '%v' passed to 'wait'", rhs.Type())
		}
		return types.Bool(reverse2.ReverseCheck(reverse, timeout))
	},
}
