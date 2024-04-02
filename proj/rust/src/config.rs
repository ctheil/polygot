use std::path::PathBuf;

use anyhow::{anyhow, Context, Result};

use crate::opts::Opts;

#[derive(Debug)]
pub struct Config {
    pub operation: Operation,
    pub pwd: PathBuf,
    pub config: PathBuf,
}

impl TryFrom<Opts> for Config {
    type Error = anyhow::Error;

    fn try_from(value: Opts) -> Result<Self> {
        let operation = value.args.try_into()?;
        let config = get_config(value.config)?;
        let pwd = get_pwd(value.pwd)?;
        Ok(Config {
            operation,
            config,
            pwd,
        })
    }
}
#[derive(Debug, PartialEq)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String),
    Remove(String),
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;

    fn try_from(value: Vec<String>) -> Result<Self> {
        let mut value = value;
        if value.len() == 0 {
            return Ok(Operation::Print(None));
        }

        let term = value.get(0).expect("expect to exist");

        if term == "add" {
            if value.len() != 3 {
                return Err(anyhow!(
                    "opeation add expects 2 args but got {}",
                    value.len() - 1
                ));
            }
            let mut drain = value.drain(1..=2);
            return Ok(Operation::Add(
                drain.next().expect("to exists"),
                drain.next().expect("to exists"),
            ));
        }
        if term == "remove" {
            if value.len() != 2 {
                return Err(anyhow!(
                    "opeation remove expects 1 args but got {}",
                    value.len() - 1
                ));
            }
            let arg = value.pop().expect("to exist");

            return Ok(Operation::Remove(arg));
        }

        if value.len() > 1 {
            return Err(anyhow!(
                "opeation print expects 0 or 1 args but got {}",
                value.len() - 1
            ));
        }
        let arg = value.pop().expect("to exist");

        Ok(Operation::Print(Some(arg)))
    }
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = config {
        return Ok(v);
    }

    let loc = std::env::home_dir().context("unable to get Application Support")?;
    let mut loc = PathBuf::from(loc);

    loc.push("projector");
    loc.push("projector.json");

    Ok(loc)
}
fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = pwd {
        return Ok(v);
    }

    Ok(std::env::current_dir().context("errored getting current directory")?)
}
#[cfg(test)]
mod test {
    use anyhow::Result;

    use super::Config;
    use crate::{config::Operation, opts::Opts};

    #[test]
    fn test_print_all() -> Result<()> {
        let opts: Config = Opts {
            args: vec![],
            config: None,
            pwd: None,
        }
        .try_into()?;
        assert_eq!(opts.operation, Operation::Print(None));
        Ok(())
    }
    #[test]
    fn test_print_key() -> Result<()> {
        let opts: Config = Opts {
            args: vec!["foo".to_string()],
            config: None,
            pwd: None,
        }
        .try_into()?;
        assert_eq!(opts.operation, Operation::Print(Some("foo".to_string())));
        Ok(())
    }
    #[test]
    fn test_add_value() -> Result<()> {
        let opts: Config = Opts {
            args: vec!["add".to_string(), "foo".to_string(), "bar".to_string()],
            config: None,
            pwd: None,
        }
        .try_into()?;
        assert_eq!(
            opts.operation,
            Operation::Add("foo".to_string(), "bar".to_string())
        );
        Ok(())
    }
    #[test]
    fn test_remove_key() -> Result<()> {
        let opts: Config = Opts {
            args: vec!["remove".to_string(), "foo".to_string()],
            config: None,
            pwd: None,
        }
        .try_into()?;
        assert_eq!(opts.operation, Operation::Remove("foo".to_string()));
        Ok(())
    }
}
