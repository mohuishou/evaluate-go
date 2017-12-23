import Env from './env';

let domain = Env == "production" ? "eva.scuplus.com" : "localhost:5000"

let config = {
    env: Env,
    domain: domain
};
export default config;