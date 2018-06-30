class Student extends Person {
    /*	
    *   Class Constructor
    *   
    *   @param firstName - A string denoting the Person's first name.
    *   @param lastName - A string denoting the Person's last name.
    *   @param id - An integer denoting the Person's ID number.
    *   @param scores - An array of integers denoting the Person's test scores.
    */
    // Write your constructor here
    constructor(firstName, lastName, id, scores){
      super(firstName, lastName, id);
      this.scores = scores;
    }
    /*	
    *   Method Name: calculate
    *   @return A character denoting the grade.
    */
    // Write your method here
    calculate(){
        // Add ints in array and take average
        var scoresAvg = (this.scores.reduce(function(a, b) { return a + b; })) / this.scores.length;

        // Return letter result based on score
        if(scoresAvg < 40){
            return 'T';
        } else if (scoresAvg >= 40 && scoresAvg < 55){
            return 'D';
        } else if (scoresAvg >= 55 && scoresAvg < 70){
            return 'P';
        } else if (scoresAvg >= 70 && scoresAvg < 80){
            return 'A';
        } else if (scoresAvg >= 80 && scoresAvg < 90){
            return 'E';
        } else if (scoresAvg >= 90 && scoresAvg <= 100){
            return 'O';
        }
    }
}
