import getConfig, { Operation } from "./config";
import getOpts from "./opts";
import Projector from "./projector";

const opts = getOpts();

const config = getConfig(opts)

const proj = Projector.fromConfig(config);

if (config.operation === Operation.Print) {
  if (config.args.length === 0) {
    console.log(JSON.stringify(proj.get_value_all()))
  } else {
    const value = proj.get_value(config.args[0]);
    if (value) {
      console.log(value)
    }
  }
}

if (config.operation === Operation.Add) {
  proj.set_value(config.args[0], config.args[1]);
  proj.save()
}

if (config.operation === Operation.Remove) {
  proj.remove_value(config.args[0]);
  proj.save()
} 
