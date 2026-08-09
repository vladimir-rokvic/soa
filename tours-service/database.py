from pymongo import MongoClient

DATABASE_URL = "mongodb://mongodb:27017"
client = MongoClient(DATABASE_URL)
db = client["project_db"]
tours_collection = db["tours"]
