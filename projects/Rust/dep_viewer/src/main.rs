mod custom_error;
use custom_error::*;

mod display;
mod fetch;
mod process;
use display::*;
use fetch::*;
use process::*;

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
    env_logger::init();

    let args = Args::parse();

    let result = process(args.path)?;
    let result = fetch(result)?;
    display_table(result);

    Ok(())
}
