import bcrypt
from datetime import datetime

from pymongo import MongoClient, collection
from pymongo.errors import CollectionInvalid
from bcrypt import hashpw

# 커스텀
from config import MONGODB_COLLECTIONS as COLS

# 컬렉션 명
USERS       = COLS['users']
FOLLOW      = COLS['follow']
ARTICLE     = COLS['article']
CONFIG      = COLS['config']
CREDENTIAL  = COLS['credential']

class DataBase(MongoClient):
    """MongoDB 설정해주는 클래스
    Args:
            host (str, optional): DB 호스트 명
            db (str, optional): DB 이름
    """
    def __init__(self, host: str, db: str):
        super().__init__(host)

        self.db = self[db]
        self.host = host         # DB 호스트 명
        self.cols = self.db.list_collection_names()

    def get_collection(self, name) -> collection.Collection:
        """데이터베이스에서 원하는 pymongo collection을 가져온다
        Args:
                name (str): COLLECTIONS 배열안에 정의된 이름만 가능
        Raises:
                pymongo.erros.CollectionInvalid: 잘못된 Collection 이름
        """
        name = name.lower()

        # Wrapper 클래스가 있으면 Wrapper로 반환
        if name in self.cols:
            return self.db[name]
        else:
            raise CollectionInvalid("콜렉션{" + name + "} 이름이 잘못되었습니다")

    def add_new_user(self, data: dict):
        """
        'users' 컬렉션에 새로운 유저를 추가한다.
            :param data: name, email. password, profile로 구성된 dict
            :return: boolean
        """
        col_user = self.get_collection(USERS)
        col_cred = self.get_collection(CREDENTIAL)

        password = data['password']
        salt = bcrypt.gensalt()

        new_user = {
            'name': data['name'],
            'email': data['email'],
            'profile': data['profile']
        }
        new_cred = {
            'email': data['email'],
            'password': hashpw(password.encode(encoding='utf-8'), salt), # salting(임의의 문자열 추가)
            'salt': salt
        }

        try:
            col_user.insert_one(new_user)
            col_cred.insert_one(new_cred)
        except:
            return False

        return True

    def get_user(self, email: str):
        """ DB에 저장된 유저 정보를 모두 가져온다. 없으면 None

        :param email: 찾고자 하는 유저의 이메일 주소
        :return: dictionary
        """
        collection = self.get_collection(USERS)

        my_query = {'email': email}

        return collection.find_one(my_query)

    def get_cred(self, email: str):
        """ DB에 저장된 유저의 보안 정보를 가져온다. 없으면 None

        :param email: 찾고자 하는 유저의 이메일 주소
        :return: dictionary
        """
        collection = self.get_collection(CREDENTIAL)

        my_query = {'email': email}

        return collection.find_one(my_query)

    def create_article(self, data: dict):
        """ 새로운 글을 등록한다.

        :param data: email, content로 이루어진 dict
        :return: None
        """
        collection = self.get_collection(ARTICLE)
        print(data)
        email = data['email']
        content = data['content']

        new_article = {
            'email': email,
            'content': content,
            'created': datetime.now(),
            'edited': datetime.now()
        }

        return collection.insert_one(new_article)