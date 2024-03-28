import { Operation } from "../config"
import Projector from "../projector"

const createData = () => {
  return {
    projector: {
      "/": {
        "foo": "bar1",
        "fem": "is_great"
      },
      "/foo": {
        "foo": "bar2"
      },
      "/foo/bar": {
        "foo": "bar3"
      }
    }
  }
}

const getProjector = (pwd: string, data = createData()): Projector => {
  return new Projector({
    args: [],
    operation: Operation.Print,
    pwd,
    config: "Hello world!"
  }, createData())
}

test("get_all", function() {
  const projector = getProjector("/foo/bar");
  expect(projector.get_value_all()).toEqual({
    "fem": "is_great",
    "foo": "bar3"
  })
})

test("get_value", function() {
  let proj = getProjector("/foo/bar");
  expect(proj.get_value("foo")).toEqual("bar3")
  proj = getProjector("/foo");
  expect(proj.get_value("foo")).toEqual("bar2")
  expect(proj.get_value("fem")).toEqual("is_great")
})

test("set_value", function() {
  let data = createData();
  let proj = getProjector("/foo/bar", data);
  proj.set_value("foo", "baz")

  expect(proj.get_value("foo")).toEqual("baz")
  proj.set_value("fem", "is_better_than_great")
  expect(proj.get_value("fem")).toEqual("is_better_than_great")

  proj = getProjector("/", data);
  expect(proj.get_value("fem")).toEqual("is_great")
})

test("remove_value", function() {
  const proj = getProjector("/foo/bar");
  proj.remove_value("fem")
  expect(proj.get_value("fem")).toEqual("is_great")
  proj.remove_value("foo")
  expect(proj.get_value("foo")).toEqual("bar2")
})

