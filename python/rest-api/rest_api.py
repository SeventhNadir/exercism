import falcon

class RestAPI(object):
    def get_users(self, req, resp):
        pass
    
    def post_create_user(self, req, resp):
        pass

    def post_create_iou(self, req, resp):
        pass


app = falcon.API()
api = RestAPI()

app.add_route('/', api)

{"users": < List of all User objects > }
{"users": < List of User objects for < users > (sorted by name)}
