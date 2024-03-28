import getConfig from "./config";
import getOpts from "./opts";

const opts = getOpts();
const config = getConfig(opts)
console.log(config)
