from os.path import isfile

"""
 from_pyfile()로 읽어들일 Flask config 파일
  * 변수명은 반드시 대문자로 설정해야 한다.
  * 공통 접두사는 Flask에서 namespace로 활용한다.
"""

# 사전검사
if not isfile('db_host.0'):
    raise Exception('db_host.0 파일이 존재하지 않습니다.')


# MongoDB Settings
with open('db_host.0') as file:
    MONGODB_HOST = file.readline()
MONGODB_DB = 'miniter'
MONGODB_COLLECTIONS = {
    'users': 'users',
    'article': 'article',
    'follow': 'follow'
}

# Custom Settings
ADMINS = [
    'daily142857@gmail.com'
]