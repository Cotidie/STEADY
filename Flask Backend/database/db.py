import bcrypt
from flask  import current_app as app

from pymongo import MongoClient, collection
from pymongo.errors import CollectionInvalid
from bcrypt import hashpw

# 커스텀
from config import MONGODB_COLLECTIONS as COLS

# 컬렉션 명
USERS   = COLS['users']
FOLLOW  = COLS['follow']
ARTICLE = COLS['article']

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
        """
        collection = self.get_collection(USERS)

        new_user = {
            'name': data['name'],
            'password': hashpw(data['password'].encode(encoding='utf-8')
                               , bcrypt.gensalt()).decode('utf-8'), # salting(임의의 문자열 추가)
            'email': data['email'],
            'profile': data['profile']
        }

        return collection.insert_one(new_user)
