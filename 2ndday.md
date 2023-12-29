- 여러 커피 값을 randomly 돌려서 coffeeshop이 돌아가게 한다.
- main.go 에서만 돌던 코드를 각 기능별로 package로 나누었다.
- receipe 가 espresso 만 있었고 이것이 orders.go 에 makeCoffee function으로 존재하고 있었는데 
  이를 recipes.go 로 따로 빼서 recipe를 늘릴 수 있는 구조, 해당 recipe를 가져다 쓸 수 있는 구조로 변경한다.
