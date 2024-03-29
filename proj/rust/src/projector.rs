use std::collections::HashMap;

use crate::config::Config;

pub struct Data {
    projector: HashMap<String, String>,
}

pub struct Projector {
    data: Data,
    config: Config,
}

impl Projector {
    pub fn set_value(&self, key: String, value: String) {
        let pwd = self.config.pwd;
        self.data.projector[&pwd].insert(key, value);
    }
}
