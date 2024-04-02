use std::{collections::HashMap, path::PathBuf};

use serde::{Deserialize, Serialize};

use crate::config::Config;

#[derive(Debug, Serialize, Deserialize)]
struct Data {
    projector: HashMap<PathBuf, HashMap<String, String>>,
}

pub struct Projector {
    data: Data,
    config: Config,
}

fn default_data() -> Data {
    Data {
        projector: HashMap::default(),
    }
}
impl Projector {
    pub fn get_value(&self, key: &str) -> Option<&String> {
        let mut curr = Some(self.config.pwd.as_path());
        let mut out = None;
        while let Some(p) = curr {
            if let Some(dir) = self.data.projector.get(p) {
                if let Some(value) = dir.get(key) {
                    out = Some(value);
                    break;
                }
            }
            curr = p.parent()
        }
        out
    }
    pub fn get_value_all(&self) -> HashMap<&String, &String> {
        let mut paths = vec![];
        let mut curr = Some(self.config.pwd.as_path());

        while let Some(p) = curr {
            paths.push(p);
            curr = p.parent();
        }
        let mut out = HashMap::new();
        for path in paths.into_iter().rev() {
            if let Some(map) = self.data.projector.get(path) {
                out.extend(map.iter());
            }
        }

        out
    }

    pub fn set_value(&mut self, key: &str, value: &str) {
        self.data
            .projector
            .get_mut(&self.config.pwd)
            .map(|x| x.insert(key.to_string(), value.to_string()));
    }
    pub fn remove_value(&mut self, key: &str) {
        self.data
            .projector
            .get_mut(&self.config.pwd)
            .map(|x| x.remove(key));
    }
    pub fn from_config(config: Config) -> Self {
        if std::fs::metadata(&config.config).is_ok() {
            let contents = std::fs::read_to_string(&config.config);
            let contents = contents.unwrap_or("{\"projector}\":{}}".to_string());
            let data = serde_json::from_str(&contents);
            let data = data.unwrap_or(default_data());

            return Projector { data, config };
        }
        Projector {
            data: default_data(),
            config,
        }
    }
}

#[cfg(test)]
mod test {

    use std::{collections::HashMap, path::PathBuf};

    use collection_macros::hashmap;

    use crate::{config::Config, projector::Projector};

    use super::Data;

    fn get_data() -> HashMap<PathBuf, HashMap<String, String>> {
        hashmap! {
            PathBuf::from("/") => hashmap! {
              "foo".to_string()=> "bar1".to_string(),
              "fem".to_string()=> "is_great".to_string()
        },
            PathBuf::from("/foo") => hashmap!{
              "foo".to_string()=> "bar2".to_string()
            },
            PathBuf::from("/foo/bar") => hashmap!{
              "foo".to_string()=> "bar3".to_string()
            },
        }
    }

    fn get_projector(pwd: PathBuf) -> Projector {
        Projector {
            config: Config {
                pwd,
                config: PathBuf::from(""),
                operation: crate::config::Operation::Print(None),
            },
            data: Data {
                projector: get_data(),
            },
        }
    }

    #[test]
    fn get_value() {
        let mut p = get_projector(PathBuf::from("/foo/bar"));
        assert_eq!(p.get_value("foo"), Some(&"bar3".to_string()));
        assert_eq!(p.get_value("fem"), Some(&"is_great".to_string()));
        p = get_projector(PathBuf::from("/foo"));
        assert_eq!(p.get_value("foo"), Some(&"bar2".to_string()));
    }
    #[test]
    fn set_value() {
        let mut p = get_projector(PathBuf::from("/foo/bar"));
        assert_eq!(p.get_value("foo"), Some(&"bar3".to_string()));
        p.set_value("foo", "baz");
        assert_eq!(p.get_value("foo"), Some(&"baz".to_string()));
    }
    #[test]
    fn remove_value() {
        let mut p = get_projector(PathBuf::from("/foo/bar"));
        p.remove_value("foo");
        p.remove_value("fem");
        assert_eq!(p.get_value("fem"), Some(&"is_great".to_string()));
        assert_eq!(p.get_value("foo"), Some(&"bar2".to_string()));
    }
    //BUG: Wont work. Expects HM<&String> but when I try &String::from(...) I get a lifetime error
    //
    // #[test]
    // fn get_value_all() {
    //     let p = get_projector(PathBuf::from("/foo/bar"));
    //     // let expected = hashmap! {
    //     //     &String::from( "foo" ) => &String::from("bar1"),
    //     //     &String::from( "fem" ) => &String::from("is_great"),
    //     //     &String::from( "foo" )=> &String::from("bar2"),
    //     //     &String::from( "foo" ) => &String::from("bar3"),
    //     // };
    //     println!("{:?}", p.get_value_all())
    //     // assert_eq!(p.get_value_all(), expected);
    // }
}
