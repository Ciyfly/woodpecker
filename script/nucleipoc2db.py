#!/usr/bin/python
# coding=utf-8
'''
Date: 2022-04-19 10:53:11
LastEditors: recar
LastEditTime: 2022-04-27 16:35:17
'''
import sqlite3
import yaml
import datetime
import sys
import os

def load_poc(poc_path):
    poc_path_list = list()
    if os.path.isdir(poc_path):
        for root, dirs, files in os.walk(poc_path):  
            for filename in files:
                suffix = os.path.splitext(filename)[1]
                if suffix == ".yml" or suffix == ".yaml":
                    poc_path_list.append(os.path.join(root, filename))
    else:
        poc_path_list.append(poc_path)
    return poc_path_list

def parse_poc(poc_path_list):
    poc_list = list()
    for poc_path in poc_path_list:
        with open(poc_path, 'r') as f:
            poc_json =  yaml.load(f)
        with open(poc_path, 'r') as t:
            poc_text =  t.read()
        poc_list.append({
            "json": poc_json,
            "text": poc_text,
        })
    return poc_list

def insert_db(poc_list, db_path):
    try:
        conn = sqlite3.connect(db_path)
        conn.text_factory=str
    except:
        print("数据库打开失败: {0}".format(db_path))
    for poc in poc_list:
        poc_json = poc.get('json')
        poc_text = poc.get('text')
        poc_name = poc_json.get('id')
        insert_sql = "INSERT INTO pocs (poc_name, content, enable, create_time, update_time, source)" \
         " values (?, ?, ?, ?, ?, ?)"
        conn.execute(insert_sql, (poc_name, poc_text, 1, datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S"), datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S"), "nuclei"))
        print("insert: {0}".format(poc_name))
    conn.commit()
    conn.close()

def main():
    print("python nucleipoc2db.py poc_path db_path")
    base_path = os.path.dirname(os.path.abspath(__file__))
    poc_path = os.path.join(base_path, "../", "pocs", "nuclei")
    db_path = os.path.join(base_path, "../", "woodpecker.db")
    if len(sys.argv)>1:
        poc_path = sys.argv[1]
    if len(sys.argv)>2:
        db_path = sys.argv[2]
    
    poc_path_list = load_poc(poc_path)
    # 还要解析yaml 获取name一些poc的信息 然后入库
    poc_list = parse_poc(poc_path_list)
    insert_db(poc_list, db_path)

if __name__ == '__main__':
    main()