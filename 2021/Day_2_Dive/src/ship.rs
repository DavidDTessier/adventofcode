trait Movement {
    fn forward(&self, position: i32);
    fn down(&self, depth: i32);
    fn up(&self, depth: i32)
}

trait Position {
    fn get_forward_position(&self) -> i32;
    fn get_depth(&self) -> i32;
}

#[derive(Debug)]
pub struct Ship {
    forward_postion: i32 = 0,
    depth: i32 = 0,
}

impl Movement for Ship {
    fn forward(&self, position: i32) {
        self.forward_postion += position;
    }
    fn dive(&self, depth: i32) {
        self.depth += depth
    }
    fn up(&self, depth: i32) {
        self.depth -= depth
    }
}

impl Position for Ship {
    fn get_depth(&self) -> i32 {
        return self.depth
    } 

    fn get_forward_position(&self) -> i32 {
        return self.forward_postion;
    }
    
}