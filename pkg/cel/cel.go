package cel

import (
	"fmt"
	"strings"
	"woodpecker/pkg/cel/proto"
	"woodpecker/pkg/cel/reverse"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
	"woodpecker/pkg/util"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types/ref"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"gopkg.in/yaml.v2"
)

type CustomLib struct {
	// 声明
	envOptions []cel.EnvOption
	// 实现
	programOptions []cel.ProgramOption
}

//	如果有set：追加set变量到 cel options
func (c *CustomLib) AddRuleSetOptions(args []yaml.MapItem) {
	for _, arg := range args {
		// 在执行之前是不知道变量的类型的，所以统一声明为字符型
		// 所以randomInt虽然返回的是int型，在运算中却被当作字符型进行计算，需要重载string_*_string
		k := arg.Key.(string)
		v := arg.Value.(string)

		var d *exprpb.Decl
		if strings.HasPrefix(v, "randomInt") {
			d = decls.NewVar(k, decls.Int)
		} else if strings.HasPrefix(v, "newReverse") {
			d = decls.NewVar(k, decls.NewObjectType("proto.Reverse"))
		} else {
			d = decls.NewVar(k, decls.String)
		}
		c.envOptions = append(c.envOptions, cel.Declarations(d))
	}
}

///////////////////////////
type CelController struct {
	Env      *cel.Env               // cel env
	ParamMap map[string]interface{} // 注入到cel中的变量
}

func InitCelOptions() CustomLib {
	custom := CustomLib{}
	custom.envOptions = []cel.EnvOption{
		cel.Container("proto"),
		//	类型注入
		cel.Types(
			&proto.UrlType{},
			&proto.Request{},
			&proto.Response{},
			&proto.Reverse{},
		),
		// 定义变量变量
		cel.Declarations(
			decls.NewVar("request", decls.NewObjectType("proto.Request")),
			decls.NewVar("response", decls.NewObjectType("proto.Response")),
		),
		// 定义
		cel.Declarations(
			bcontainsDec, iContainsDec, bmatchDec, md5Dec,
			//startsWithDec, endsWithDec,
			inDec, randomIntDec, randomLowercaseDec,
			base64StringDec, base64BytesDec, base64DecodeStringDec, base64DecodeBytesDec,
			urlencodeStringDec, urlencodeBytesDec, urldecodeStringDec, urldecodeBytesDec,
			substrDec, sleepDec, reverseWaitDec,
		),
	}
	// 实现
	custom.programOptions = []cel.ProgramOption{cel.Functions(
		containsFunc, iContainsFunc, bcontainsFunc, matchFunc, bmatchFunc, md5Func,
		//startsWithFunc,  endsWithFunc,
		inFunc, randomIntFunc, randomLowercaseFunc,
		base64StringFunc, base64BytesFunc, base64DecodeStringFunc, base64DecodeBytesFunc,
		urlencodeStringFunc, urlencodeBytesFunc, urldecodeStringFunc, urldecodeBytesFunc,
		substrFunc, sleepFunc, reverseWaitFunc,
	)}
	return custom
}

func (c *CustomLib) CompileOptions() []cel.EnvOption {
	return c.envOptions
}

func (c *CustomLib) ProgramOptions() []cel.ProgramOption {
	return c.programOptions
}

// 初始化env
func (celController *CelController) InitCel(poc *parse.Poc) {
	//	1.生成cel env环境
	option := InitCelOptions()
	//	注入set定义的变量
	if poc.Set != nil {
		option.AddRuleSetOptions(poc.Set)
	}
	env, err := cel.NewEnv(cel.Lib(&option))
	if err != nil {
		log.Errorf("init cel env poc: %s error: %s", poc.Name, err)
	}
	celController.Env = env
	celController.ParamMap = make(map[string]interface{})
}

// 处理poc: set
func (cc *CelController) InitSet(poc *parse.Poc, newReq *proto.Request) (err error) {
	// 如果没有set 就直接返回
	if len(poc.Set) == 0 {
		return
	}
	cc.ParamMap["request"] = newReq
	for _, setItem := range poc.Set {
		key := setItem.Key.(string)
		value := setItem.Value.(string)
		// 反连平台
		if value == "newReverse()" {
			cc.ParamMap[key] = reverse.NewReverse()
			continue
		}
		out, err := Evaluate(cc.Env, value, cc.ParamMap)
		if err != nil {
			return err
		}
		switch value := out.Value().(type) {
		// set value 无论是什么类型都先转成string
		case *proto.UrlType:
			cc.ParamMap[key] = util.UrlTypeToString(value)
		case int64:
			cc.ParamMap[key] = int(value)
		default:
			cc.ParamMap[key] = fmt.Sprintf("%v", out)
		}
	}
	return
}

//	计算单个表达式
func Evaluate(env *cel.Env, expression string, params map[string]interface{}) (ref.Val, error) {
	ast, iss := env.Compile(expression)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	prg, err := env.Program(ast)
	if err != nil {
		return nil, err
	}
	out, _, err := prg.Eval(params)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// 计算cel表达式
func (cc *CelController) Evaluate(char string) (bool, error) {
	out, err := Evaluate(cc.Env, char, cc.ParamMap)
	if err != nil {
		log.Errorf("cel Evaluate error", err)
		return false, err
	}
	if fmt.Sprintf("%v", out) == "false" {
		return false, nil
	}
	return true, nil
}
