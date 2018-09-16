# Growth Book API

[[https://github.com/azbshiri/growth-book/blob/master/public/growth.png|alt=growth]]
An web implementation of the growth book, designed by John Fish at Harvard, is a daily planner built to encourage a growth mindset and flow in all activities. 


#### Daily Scheduling
Writing out a daily schedule ensures that you stay committed to your goals for every hour of the day and that nothing slips your mind. Keeping a daily schedule keeps you accountable and productive.

#### To Do Listing
Every day has tasks that need to be done. Keeping track of these tasks helps you be more productive.
Goal Setting
An important part of being engaged in a flow activity is having goals. Creating goals for every day allows you to treat each day as a flow experience, therefore increasing your happiness and enjoyment in life.

#### Motivational Awareness
What is the big, driving force that is motivating you to achieve your goals? Being aware of this is what helps promote harmony in your life, making it possible to reflect on a series of days with joy because you followed an overarching theme throughout them.

#### Happiness Reflection
Gratitude journaling has been shown time and time again to increase happiness, productivity, and general fulfillment in life. Keeping track of things that you are grateful for and that made you happy will make your life better.


# Run

```
./runner.sh
```

# Dependences
You can jump into this step if you want to know what dependences are used and install them yourself.

```
go get -v github.com/azbshiri/common/db
go get -v github.com/go-pg/pg
go get -v github.com/joho/godotenv
go get -v github.com/labstack/echo
```


# To-Do
- [x] Add Echo intergration
- [x] Use PostgreSQL as database
- [x] Add database migrations (using `dbmate`)
- [x] Use Repository pattern to access database layer
- [ ] Add authentication mechanism
- [ ] Dockerize application
- [ ] Add unit and integeration tests
