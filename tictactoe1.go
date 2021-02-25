package main

import("fmt")

func main(){
  gameOver := false //bien tro choi ket thuc
  board := [9]int{0,0,0,0,0,0,0,0,0}
  turnNumber := 1

  for gameOver != true{
    presentBoard(board)
    player := 0
    if turnNumber % 2 == 1{
      fmt.Println("Day la luot cua X")
      player = 1
    }else{
      fmt.Println("Day la luot cua O")
      player = 2
    }
  
    currentMove := askForPlay() //dung lai voi 9
      if currentMove == 9 {
      return
      }
    board = executePM(currentMove, player ,board)
    result := checkForWinner(board)
    if result > 0{
      fmt.Printf("Nguoi choi %d da thang cuoc \n\n", result)
      gameOver = true
    }else if turnNumber == 9{
      fmt.Println("Game nay hoa!\n\n")
      gameOver = true
    }else {
      turnNumber++
    }
  }
}
func presentBoard(b [9]int){
  for i, v := range b {
    if v == 0{ //Khong gian trong de cho so vao
      fmt.Printf("%d",i)
    }else if v == 1{
      fmt.Printf("X")
    }else if v == 10{
      fmt.Printf("O")
    }
    if i>0 && (i+1)%3==0{
      fmt.Printf("\n")
    }else{
      fmt.Printf("|")
    }
  }
}

func askForPlay()int{
  fmt.Println("Nguoi X di truoc")
  var moveInt int
  fmt.Scan(&moveInt)
  return moveInt
}

func executePM(moveInt int,player int, b [9]int)[9]int{
  if b[moveInt] != 0{
    fmt.Println("Chon nuoc con trong.")
    moveInt = askForPlay()
    b = executePM(moveInt, player, b)
  }else{
    if player == 1{
      b[moveInt] = 1
    }else if player == 2{
      b[moveInt] = 10
    }
  }
  for moveInt > 9{
    fmt.Println("Vui long chon 1 so nho hon 10.")
    moveInt = askForPlay()
  }
  if player == 1{
    b[moveInt] = 1
  }else if player == 2{
    b[moveInt] = 10
  }
  return b
}

func checkForWinner(b [9]int) int {
  sums := [8] int {0,0,0,0,0,0,0,0}

  sums[0] = b[2]+b[4]+b[6]
  sums[1] = b[0]+b[3]+b[6]
  sums[2] = b[1]+b[4]+b[7]
  sums[3] = b[2]+b[5]+b[8]
  sums[4] = b[0]+b[4]+b[8]
  sums[5] = b[6]+b[7]+b[8]
  sums[6] = b[3]+b[4]+b[5]
  sums[7] = b[0]+b[1]+b[2]
  for _, v := range sums {
    if v == 3{
      return 1
    } else if v == 30{
      return 2
    }
  }
  return 0
}

