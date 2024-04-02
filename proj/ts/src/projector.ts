import * as path from "path"
import { Config } from "./config"
import * as fs from "fs"

export type Data = {
  projector: {
    [key: string]: {
      [key: string]: string
    }
  }
}


const defaultData = {
  projector: {}
}
export default class Projector {

  constructor(private c: Config, private data: Data) { }

  get_value_all(): { [key: string]: string } {
    let curr = this.c.pwd
    let prev = ""
    const paths = [];
    do {
      prev = curr;
      paths.push(curr);
      curr = path.dirname(curr)
    } while (curr !== prev);
    return paths.reverse().reduce((acc, path) => {
      const value = this.data.projector[path];
      if (value) {
        Object.assign(acc, value)
      }
      return acc;
    }, {})
  }

  get_value(key: string): string | undefined {
    let curr = this.c.pwd
    let prev = ""
    let out: string | undefined = undefined
    do {
      const value = this.data.projector[curr]?.[key];
      if (value) {
        out = value;
        break
      }
      prev = curr;
      curr = path.dirname(curr)
    } while (curr !== prev)
    return out;
  }

  set_value(key: string, value: string) {
    let pwd = this.c.pwd;
    if (!this.data.projector[pwd]) {
      this.data.projector[pwd] = {}
    }
    this.data.projector[pwd][key] = value
  }


  remove_value(key: string) {
    const dir = this.data.projector[this.c.pwd];
    if (dir) {
      delete dir[key]
    }
  }

  save() {
    const p = path.dirname(this.c.config);
    if (!fs.existsSync(p)) {
      fs.mkdirSync(p, { recursive: true })
    }
    fs.writeFileSync(this.c.config, JSON.stringify(this.data));
  }

  static fromConfig(c: Config): Projector {
    if (fs.existsSync(c.config)) {
      let data: Data = defaultData;
      try {
        data = JSON.parse(
          fs.readFileSync(c.config).toString());
      } catch (e) {
        console.warn("config does not exist... using default data")
      }
      return new Projector(c, data)
    }
    return new Projector(c, defaultData)
  }



}
