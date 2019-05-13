import tornado.ioloop
import tornado.web
import tornado.autoreload
import os

application = tornado.web.Application([
    (r"/()$", tornado.web.StaticFileHandler, {'path':'./index.html'}),
    (r"/(.*)", tornado.web.StaticFileHandler, {"path": "."}),
],debug=True)

if __name__ == "__main__":
    tornado.autoreload.watch(os.path.abspath('.'))
    tornado.autoreload.watch(os.path.abspath('./lib'))
    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()