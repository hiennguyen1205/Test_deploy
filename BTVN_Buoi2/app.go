package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
	Gender string `json:"gender"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("personsmall.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	fmt.Println("====2.1) Thống kê người theo thành phố====")
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}
	fmt.Println()

	fmt.Println("====2.2) Số lượng người trong mỗi công việc====")
	peopleByJob := GroupPeopleByJob(people)
	for key, value := range peopleByJob {
		fmt.Print(key + " - ")
		fmt.Println(value)
	}
	fmt.Println()

	fmt.Println("====2.3) Top 5 job được nhiều người làm nhất====")
	top5Job := Top5JobsByNumer(peopleByJob)
	for _, job := range top5Job {
		fmt.Print(job.NameJob + " - ")
		fmt.Println(job.Number)
	}
	fmt.Println()

	fmt.Println("====2.4) Top 5 thành phố nhiều người ở nhất ====")
	top5Cities := Top5CitiesByNumber(people)
	for _, city := range top5Cities {
		fmt.Print(city.NameCity + " : ")
		fmt.Println(city.NumberPeople)
	}
	fmt.Println()

	fmt.Println("====2.5) Top công việc trong thành phố ====")
	topJobOnCity := TopJobByNumerInEachCity(people)
	for key, value := range topJobOnCity {
		fmt.Print(key + " : ")
		fmt.Println(value)
	}
	fmt.Println()

	fmt.Println("====2.6) Lương trung bình mỗi nghề ====")
	averageSalary := AverageSalaryByJob(people)
	for key, value := range averageSalary {
		fmt.Print(key + " : ")
		fmt.Println(value)
	}
	fmt.Println()

	fmt.Println("====2.7) 5 thành phố có mức lương trung bình cao nhất ====")
	averageSalaryInCity := FiveCitiesHasTopAverageSalary(people)
	for i := 0; i < 5; i++ {
		fmt.Print("Thành phố: " + averageSalaryInCity[i].NameCity)
		fmt.Print(", Mức lương: ")
		fmt.Println(averageSalaryInCity[i].AverageSalary)
	}
	fmt.Println()

	fmt.Println("====2.8) 5 thành phố có mức lương Developer cao nhất ====")
	averageSalaryDevInCity := FiveCitiesHasTopSalaryForDeveloper(people)
	for _, value := range averageSalaryDevInCity {
		fmt.Print("Thành phố: " + value.NameCity)
		fmt.Print(", Mức lương: ")
		fmt.Println(value.AverageSalary)
	}
	fmt.Println()

	fmt.Println("====2.9) Tuổi trung bình từng nghề nghiệp ====")
	averageAgePerJob := AverageAgePerJob(people)
	for job, averageAge := range averageAgePerJob{
		fmt.Print("Job: "+ job+", Average Age: ")
		fmt.Println(averageAge)
	}
	fmt.Println()

	fmt.Println("====2.10) Tuổi trung bình từng thành phố ====")
	averageAgePerCity := AverageAgePerCity(people)
	for city, averageAge := range averageAgePerCity {
		fmt.Print("City: "+ city+", Average Age: ")
		fmt.Println(averageAge)
	}
	fmt.Println()

	fmt.Println("====2.11) Giới tính từng thành phố ====")
	NumberGenderInCity(people)
}
