mod custom_error;
use custom_error::*;

mod process;
use process::*;

mod fetch;
use fetch::*;

use clap::Parser;

#[derive(Debug, Parser)]
#[command(about = "Get information on dependencies of your projects")]
struct Args {
    #[arg(
        short,
        long,
        help = "Path to project directory or file.",
        long_help = "Specify the path to a dependency file or to a directory containing the file file. Check out the documentation for more information."
    )]
    path: String,

    #[arg(short, long, default_value = "false")]
    graph: bool,
}

fn main() -> Result<(), AkkadiaError> {
    let args = Args::parse();

    dbg!(&args);

    let result = process(args.path)?;
    let result = fetch(result)?;

    for dep in result {
        println!(
            "[{}] {} {}",
            dep.set.dep_type, dep.set.name, dep.set.version
        );
    }

    Ok(())
}
