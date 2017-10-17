<!DOCTYPE html>
<html lang="en">
  <head>

	package main
		import(
		"fmt"
		"math/rand"
		"time"
	)//this generates random number between given range
	
	func xrand(min, max int) int {
		rand.Seed(time.Now().Unix())
		return rand.Intn(max - min) + min
	}

	func main() {
		var myname string
		myrand := xrand(1, 6)
		guessTaken := 0
		var guess int

		fmt.Println("Hello! What is your name?")
		fmt.Scanf("%s", &myname)
		fmt.Printf("Well, %s, I am thinking of a number between 1 and 6.\n", myname)

				<!-- Required meta tags -->
				<meta charset="utf-8">
				<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

				<!-- Bootstrap CSS -->
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
				//this is the while loop
				for guessTaken < 6 {
				fmt.Println("Take a guess...")
				fmt.Scanf("%d", &guess)
				guessTaken++
				
					if guess < myrand {
					fmt.Println("Your guess is too low.")
					}
					if guess > myrand {
					fmt.Println("Your guess is too high.")
					}
					if guess == myrand {
					break
						}
					}
					if guess == myrand {
					fmt.Printf("Good job %s! You guessed my number in %d tries\n", myname, guessTaken)
					} 
					else {
					fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
			}
		}

  </head>
  <body>
    <h1>Hello, world!</h1>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
  </body>
</html>