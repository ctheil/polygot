import cli from "command-line-args"

// get args into obj

export type Opts = {
  args?: string[],
  pwd?: string, // present working dir
  config?: string,
}

export default function getOpts(): Opts {
  return cli([{
    name: "args",
    defaultOption: true,
    type: String,
    multiple: true,
  }, {
    name: "config",
    alias: "c",
    type: String,
  }, {
    name: "pwd",
    alias: "p",
    type: String,
  }]) as Opts

}
