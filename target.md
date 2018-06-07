# Мишень

Маруся готовится к экзамену: изучает кривые второго порядка и повторяет темы попроще. Маруся решила
нарисовать сову.
Контур головы задан уравнением: 1/2 x^2 + y2 = 1, 

бровей: y = 0,5 |x| + 0,5, 

первый глаз: (x-0,5)^2 + y^2 = 0,3,

а второй глаз Маруся нарисовала так, что сова получилась симпатичной и с сохранением симметрии.

После этого она распечатала полученный рисунок и сделала перерыв: начала кидать в рисунок дротики.
Успешным попаданием считается попадание строго ниже бровей строго внутри головы, строго вне глаз. 
Напишите программу, которая по координатам определяет, успешным ли было попадание.

### Входные данные

Строка содержит x и y (|x|<=100, |y|<=100).

### Результат работы

Выведите 'YES', если попадание успешное и 'NO' в ином случае.

## Примеры

### Входные данные

    0 -0.75
    
### Результат работы

    YES
    