use hex;

use quiche::h3::NameValue;
use quiche::h3::qpack;

fn main() {
    let mut args = std::env::args();
    let cmd = &args.next().unwrap();

    if args.len() != 1 {
        println!("Usage: {cmd} <input-file>");
        return;
    }

    let inputfile = &args.next().unwrap();
    let inputhexstring = std::fs::read_to_string(inputfile).unwrap();

    let data = hex::decode(inputhexstring.trim()).unwrap();

    let mut dec = qpack::Decoder::new();

    for hdr in dec.decode(&data, u64::MAX).unwrap() {
        let name = std::str::from_utf8(hdr.name()).unwrap();
        let value = std::str::from_utf8(hdr.value()).unwrap();
        println!("key={name}\n\tvalue={value}");
    }
}
