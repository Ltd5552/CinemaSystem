package main

// QueryLeftTicket 查询所有取票机剩余可印票数
func (a *Administrator) QueryLeftTicket() {
	queryTicketMachine()
}

// QueryAllEvaluation 查询所有评论
func (a *Administrator) QueryAllEvaluation() {
	queryEvaluation()
}

// QuerySumEvaluation 查询所有用户评论总数
func (a *Administrator) QuerySumEvaluation() {
	querySumEvaluationByUid()
}

// InsertScreening 插入场次信息
func (a *Administrator) InsertScreening(screeningNum string, movieNum string, theaterNum string, showTime string, remainSeats int64) {
	insertScreenings(screeningNum, movieNum, theaterNum, showTime, remainSeats)
}

// InsertMovie 插入电影信息
func (a *Administrator) InsertMovie(movieNum string, movieTitle string, releaseDate string, duration float64, aveFilmScore float64) {
	insertMovies(movieNum, movieTitle, releaseDate, duration, aveFilmScore)
}

// DeleteMovie 删除电影信息
func (a *Administrator) DeleteMovie(name string) {
	deleteMovies(name)
}

// DeleteScreening 删除场次信息
func (a *Administrator) DeleteScreening(num string) {
	deleteScreenings(num)
}

// QueryMovie 查看电影信息
func (a *Administrator) QueryMovie(name string) {
	queryMovies(name)
}

// QueryScreening 查看所有场次信息
func (a *Administrator) QueryScreening() {
	queryScreenings()
}
