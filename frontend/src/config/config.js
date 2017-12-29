import Env from './env';

let domain = Env == "production" ? "115.159.94.101:5000" : "localhost:5000"

let config = {
    env: Env,
    domain: domain
};
export default config;