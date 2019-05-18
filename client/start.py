from subprocess import Popen, PIPE

import tornado.ioloop
import tornado.web
import tornado.autoreload
import tornado.template
import tornado.escape

import os

debug = True
page = "<h1>Internal server error. Could not find content</h1>"
config = "{ error: 1 }"


class Index(tornado.web.RequestHandler):
    def get(self):
        if debug:
            loadData()
        print(config)
        self.write(page)


def main():
    tornado.autoreload.watch(os.path.abspath('.'))
    tornado.autoreload.watch(os.path.abspath('./lib'))
    tornado.autoreload.watch(os.path.abspath('./lib/render'))
    tornado.autoreload.watch(os.path.abspath('./lib/controls'))
    tornado.autoreload.watch(os.path.abspath('./lib/server'))
    loadData()

    tornado.web.Application([
        (r"/$", Index),
        (r"/(.*)", tornado.web.StaticFileHandler, {"path": "."}),
    ],debug=debug).listen(8888)

    tornado.ioloop.IOLoop.instance().start()

def loadData():
    config = readConfig('../config')

    with open("./index.html", "r") as f:
        global page
        data = f.read()
        page = tornado.escape.squeeze(data.replace("JCONFIG", config))

def readConfig(loc):
    p = Popen(['go', 'run', 'github.com/petelliott/shooty-game/jsondir/d2json', loc],
            stdin=PIPE, stdout=PIPE, stderr=PIPE)
    output, e = p.communicate(b"input data that is passed to subprocess' stdin")
    if e:
        print(e)

    return output.decode("utf-8")

if __name__ == "__main__":
    main()
