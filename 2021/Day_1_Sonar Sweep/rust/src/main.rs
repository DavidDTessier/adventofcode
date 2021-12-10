use std::fs::File;
use std::io::{BufRead, BufReader, Error, ErrorKind};


// function takes path in the form of string slice and returns enum
// which contains vector of integers on success or IO error type, see `std::io::Error`
fn read(path: &str) -> Result<Vec<i32>, Error> {
    let file = File::open(path)?; // open file by given path
    // wrap file into generic buffered reader, it will use 4 KB buffer internally
    // to reduce number of syscalls, thus improving performance
    let br = BufReader::new(file);
    // create an empty vector, type of the stored elements will be inferred
    let mut v = Vec::<i32>::new();
    // br.lines() creates an iterator over lines in the reader
    // see: https://doc.rust-lang.org/std/io/trait.BufRead.html#method.lines
    for line in br.lines() {
        // IO operations generally can return error, we check if got
        // an error,in which case we return this error as the function result
        let line = line?;
        let n = line   
            .trim() // trim "whitespaces"
            .parse() // call `str::parse::<i64>(&self)` method on the trimmed line, which parses integer
            .map_err(|e| Error::new(ErrorKind::InvalidData, e))?; // parse() can return error (e.g. for string "abc"), here if we got it, we convert it to `std::io::Error` type and return it as function result
        v.push(n); // push acquired integer to the vector
    }
    Ok(v) // everything is Ok, return vector
}

fn process_compare_previous(data: &Vec::<i32>) {

    let mut previous: i32 = 0;
    let mut increased_count: i32 = 0;
    for (pos, current) in data.into_iter().enumerate() {
        if pos == 0 {
            previous = *current;
        } else {
            if *current > previous {
                increased_count += 1;
            }

            previous = *current;
        }
    }
    println!("Number of increases {}", increased_count);
}

fn process_sliding_window_comparison(data: &Vec::<i32>) {
    let mut previous_sum: i32 = 0;
    let mut increased_count: i32 = 0;

    let windowed_data = data.windows(3);

    for window in windowed_data {
        if previous_sum == 0 {
            previous_sum = window.iter().sum();
        } else {
            let current_sum : i32 = window.iter().sum();
            if current_sum > previous_sum {
                increased_count += 1;
            }
            previous_sum = current_sum;
        }
    }

    println!("Number of sums increases {}", increased_count);
}


// This is the main function
fn main() {
    let filename = "input/input.dat";
    let vec = read(filename);
    let data = vec.iter().next().unwrap(); // get the vector from the result

    process_compare_previous(&data);
    process_sliding_window_comparison(&data);
  
}
