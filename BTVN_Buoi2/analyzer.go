package main

import (
	"fmt"
	"sort"
	"time"
)

//2.1 Gom tất cả những người trong cùng một thành phố lại
func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

//2.2 Nhóm các nghề nghiệp và đếm số người làm
func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}


//2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp
type NumberOnJob struct {
	NameJob string
	Number  int
}

func Top5JobsByNumer(job map[string]int) []NumberOnJob {
	var listJob []NumberOnJob

	for key, value := range job {
		listJob = append(listJob, NumberOnJob{key, value})
	}

	sort.Slice(listJob, func(i, j int) bool { return listJob[i].Number > listJob[j].Number })

	return listJob[0:5]
}

//2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp
type NumberInCity struct {
	NameCity     string
	NumberPeople int
}

func Top5CitiesByNumber(p []Person) []NumberInCity {
	var listCities []NumberInCity
	numberPeopleInCity := make(map[string]int)
	for _, person := range p {
		numberPeopleInCity[person.City]++
	}
	for key, value := range numberPeopleInCity {
		listCities = append(listCities, NumberInCity{key, value})
	}
	sort.Slice(listCities, func(i, j int) bool { return listCities[i].NumberPeople > listCities[j].NumberPeople })

	return listCities[0:5]
}

//2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất
func TopJobByNumerInEachCity(p []Person) (result map[string]NumberOnJob) {
	result = make(map[string]NumberOnJob)
	numberJobsInCity := JobInCity(p)
	for key, value := range numberJobsInCity {
		result[key] = countJob(value)
	}
	return result
}

//tính số lượng công việc trong thành phố
func JobInCity(p []Person) (result map[string][]string) {
	result = make(map[string][]string)
	for _, person := range p {
		result[person.City] = append(result[person.City], person.Job)
	}
	return result
}


//2.6 Ứng với một nghề, hãy tính mức lương trung bình
//tính số lượng công việc
func countJob(listJob []string) (result NumberOnJob) {
	jobs := make(map[string]int)
	for _, job := range listJob {
		jobs[job]++
	}
	var tmp []NumberOnJob
	for key, value := range jobs {
		tmp = append(tmp, NumberOnJob{key, value})
	}
	sort.Slice(tmp, func(i, j int) bool { return tmp[i].Number > tmp[j].Number })
	return tmp[0]
}

func AverageSalaryByJob(p []Person) (result map[string]float32) {
	//tính tổng người trong 1 công việc
	totalPeopleInJob := make(map[string][]Person)
	for _, person := range p {
		totalPeopleInJob[person.Job] = append(totalPeopleInJob[person.Job], person)
	}
	result = make(map[string]float32)
	for jobName, listPerson := range totalPeopleInJob {
		totalSalary := 0
		for _, person := range listPerson {
			totalSalary += person.Salary
		}
		result[jobName] = float32(totalSalary / len(listPerson))
		// fmt.Println(totalSalary)
	}
	return result

}

//2.7 Năm thành phố có mức lương trung bình cao nhất
type AverageSalaryByCity struct {
	NameCity      string
	AverageSalary int
}

func FiveCitiesHasTopAverageSalary(p []Person) (result []AverageSalaryByCity) {
	salaryInCity := make(map[string]int)
	//tính lương của mỗi thành phố
	for _, person := range p {
		salaryInCity[person.City] += person.Salary
	}

	// tính số người mỗi thành phố
	numberPeopleInCity := make(map[string]int)
	for _, person := range p {
		numberPeopleInCity[person.City]++
	}

	//tính lương trung bình mỗi thành phố
	for city := range salaryInCity {
		salaryInCity[city] = salaryInCity[city] / numberPeopleInCity[city]
	}

	//sắp xếp lại để lấy ra 5 thành phố có mức lương cao nhất
	for nameCity, averageSalary := range salaryInCity {
		result = append(result, AverageSalaryByCity{nameCity, averageSalary})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].AverageSalary > result[j].AverageSalary })
	return result
}


//2.8 Năm thành phố có mức lương trung bình của developer cao nhất
func FiveCitiesHasTopSalaryForDeveloper(p []Person) (result []AverageSalaryByCity) {
	//tính lương của dev trong mỗi thành phố
	salaryDevInCity := make(map[string]int)
	for _, person := range p {
		if person.Job == "developer" {
			salaryDevInCity[person.City] += person.Salary
		}
	}

	//tính tổng dev trong mỗi thành phố
	numberDevInCity := make(map[string]int)
	for _, person := range p {
		if person.Job == "developer" {
			numberDevInCity[person.City]++
		}
	}

	//tính lương trung bình của dev trong mỗi thành phố
	for city := range salaryDevInCity {
		salaryDevInCity[city] = salaryDevInCity[city] / numberDevInCity[city]
	}

	//sắp xếp để lấy ra 5 thành phố có mức lương dev cao nhất
	for city, salary := range salaryDevInCity {
		result = append(result, AverageSalaryByCity{city, salary})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].AverageSalary > result[j].AverageSalary })
	if len(result) >= 5 {
		return result[0:5]
	} else {
		return result
	}
}


//2.9 Tuổi trung bình từng nghề nghiệp
func AverageAgePerJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	//tổng số tuổi trong mỗi nghề
	for _, person := range p {
		birthday := person.Birthday
		birthdayConvert, _ := time.Parse(time.RFC3339, birthday+"T00:00:00Z") //convert string to date VD: 2003-07-19 00:00:00 +0000 UTC
		result[person.Job] += time.Now().Year() - birthdayConvert.Year()
	}
	//tổng số nguồi trong mỗi nghề
	numberPeopleOnJob := make(map[string]int)
	for _, person := range p {
		numberPeopleOnJob[person.Job]++
	}
	//tính tuổi trung bình trong mỗi nghề nghiệp
	for job := range result {
		result[job] = result[job]/numberPeopleOnJob[job]
	}
	return result
}

//2.10 Tuổi trung bình ở từng thành phố
func AverageAgePerCity(p []Person) (result map[string]int) {
	// tính số người mỗi thành phố
	result = make(map[string]int)
	for _, person := range p {
		result[person.City]++
	}
	//tính tổng số tuổi mỗi thành phố
	totalAgeInCity := make(map[string]int)
	for _, person := range p {
		birthday := person.Birthday
		birthdayConvert, _ := time.Parse(time.RFC3339, birthday+"T00:00:00Z") //convert string to date VD: 2003-07-19 00:00:00 +0000 UTC
		totalAgeInCity[person.City] += time.Now().Year() - birthdayConvert.Year()
	}

	//tính tuổi trung bình mỗi thành phố
	for city := range result {
		result[city] = totalAgeInCity[city] / result[city]
	}
	return result
}
type NumberGender struct {
	male int
	female int
}
//2.11 mỗi thành phố có bao nhiêu nam nữ
func NumberGenderInCity(p []Person) (result map[string][]NumberGender) {

	result = make(map[string][]NumberGender)
	//tính số người mỗi thành phố
	peopleInCity := make(map[string][]Person)
	for _, person := range p {
		peopleInCity[person.City] = append(peopleInCity[person.City], person)
	}
	// fmt.Println(peopleInCity)
	//tổng số giới tính mỗi thành phố
	genderMale :=  0
	genderFemale :=  0

	for cityName, listPerson := range peopleInCity {
		fmt.Println(cityName)
		for _, person := range listPerson{
			if person.Gender =="Male"{
				genderMale++
			} else {
				genderFemale++
			}
		}
		fmt.Println(NumberGender{genderMale,genderFemale})
		result[cityName] = append(result[cityName], NumberGender{genderMale,genderFemale})
	}
	
	return result
}