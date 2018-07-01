function __construct($elements) {
  $this->elements = $elements;
}

# wtf is php syntax might as well use mips this took 
# too long to figure out never do this again

function computeDifference(){
  $max = max($this->elements);
  $min = min($this->elements);

  $this->maximumDifference = $max - $min;
}
