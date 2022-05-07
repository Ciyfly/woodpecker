#!/usr/bin/python
# coding=utf-8
'''
Date: 2022-04-27 16:29:36
LastEditors: recar
LastEditTime: 2022-04-27 16:49:38
'''
#!/usr/bin/python
# coding=utf-8
'''
Date: 2022-04-19 10:53:11
LastEditors: recar
LastEditTime: 2022-04-22 14:49:14
'''
import sqlite3
import yaml
import datetime
import sys
import os

def load_scripts(scripts_path):
    scripts_name_list = list()
    if os.path.isdir(scripts_path):
        for root, dirs, files in os.walk(scripts_path):  
            for filename in files:
                if filename!="scripts.go":
                    scripts_name_list.append(filename.split(".")[0])
    else:
        scripts_name_list.append(filename)
    return scripts_name_list

def insert_db(scripts_name_list, db_path):
    try:
        conn = sqlite3.connect(db_path)
        conn.text_factory=str
    except:
        print("数据库打开失败: {0}".format(db_path))
    for script_name in scripts_name_list:
        insert_sql = "INSERT INTO pocs (poc_name, enable, create_time, update_time, source)" \
         " values (?, ?, ?, ?, ?)"
        conn.execute(insert_sql, (script_name, 1, datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S"), datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S"), "script"))
        print("insert: {0}".format(script_name))
    conn.commit()
    conn.close()

def main():
    print("python goscript2db.py scripts_path db_path")
    base_path = os.path.dirname(os.path.abspath(__file__))
    scripts_path = os.path.join(base_path, "../", "pocs", "scripts")
    db_path = os.path.join(base_path, "../", "woodpecker.db")
    if len(sys.argv)>1:
        scripts_path = sys.argv[1]
    if len(sys.argv)>2:
        db_path = sys.argv[2]
    
    scripts_name_list = load_scripts(scripts_path)
    insert_db(scripts_name_list, db_path)

if __name__ == '__main__':
    main()