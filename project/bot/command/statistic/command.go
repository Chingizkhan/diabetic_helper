package statistic

import (
	"context"
	"diabetHelperTelegramBot/bot/service/grpc/diabet_helper"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"strings"
	"time"
)

func Handle(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	sl, err := diabet_helper.NewClient().FindSL(ctx, &pb.FindSLRequest{
		Pagination:      &pb.Pagination{Limit: -1},
		Value:           "",
		UserId:          c.Sender().ID,
		CreatedAtStart:  0,
		CreatedAtFinish: 0,
	})
	if err != nil {
		return err
	}

	average := CountAverage(sl)

	//var sum float64
	//res := make([]string, 0, len(sl.Data))
	var dates = make(map[string][]string, 50)

	for _, l := range sl.Data {
		t := time.Unix(l.CreatedAt, 0)
		year, month, day := t.Date()
		key := fmt.Sprintf("%s.%s.%d", addZero(day), addZeroMonth(month), year)
		dates[key] = append(dates[key], fmt.Sprintf(
			"%s \t⟞\t %s\n",
			FormatDate(l.CreatedAt),
			l.Value,
		))
	}

	// sort

	r := make([]string, 0, 10)
	for k, v := range dates {
		d := fmt.Sprintf(`

					%s
%s`, k, join(v...))
		r = append(r, d)
	}

	m := join(r...)

	//for _, l := range sl.Data {
	//	val, err := strconv.ParseFloat(l.Value, 64)
	//	if err != nil {
	//		return err
	//	}
	//	sum += val
	//	res = append(res, fmt.Sprintf(
	//		"\n %s \t⟞\t %s",
	//		FormatDate(l.CreatedAt),
	//		l.Value,
	//	))
	//}
	//sum /= float64(len(sl.Data))
	//average := strconv.FormatFloat(sum, 'f', 1, 64)
	//
	//return c.Send(fmt.Sprintf("average: %v \n\nsugar levels: %v", average, res))
	return c.Send(fmt.Sprintf("average: %s \n\nsugar levels: %v", average, m))
}

func CountAverage(sl *pb.SugarLevels) (average string) {
	var sum float64
	for _, l := range sl.Data {
		f, err := strconv.ParseFloat(l.Value, 64)
		if err != nil {
			log.Println("err CountAverage:", err)
			return ""
		}
		sum += f
	}
	sum /= float64(len(sl.Data))
	return strconv.FormatFloat(sum, 'f', 1, 64)
}

func FormatDate(date int64) string {
	loc := time.FixedZone("UTC-8", +6*60*60)
	d := time.Unix(date, 0).In(loc)

	newDate := fmt.Sprintf("%s:%s", addZero(d.Hour()), addZero(d.Minute()))

	return newDate
}

func addZero(i int) string {
	s := strconv.Itoa(i)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}

func addZeroMonth(t time.Month) string {
	newT := fmt.Sprintf("%d", t)
	if len(newT) == 1 {
		newT = "0" + newT
	}
	return newT
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
