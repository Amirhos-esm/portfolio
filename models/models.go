package models

import (
	"time"

	"github.com/Amirhos-esm/portfolio/util"
)

type DataStore struct {
	Title               string              `json:"title"`
	Skills              Skills              `json:"skills"`
	PersonalInformation PersonalInformation `json:"personal_information"`
	Experiences         []Experience        `json:"experiences"`
	Educations          []Education         `json:"educations"`
	Projects            []Project           `json:"projects"`
}

func GetMockData() DataStore {
	return DataStore{
		Title: "Portfolio – Software Engineer",
		PersonalInformation: PersonalInformation{
			FullName:          "John Doe",
			ProfessionalTitle: "Software Engineer",
			Bio:               "Software engineer experienced in building scalable systems and modern web applications.",
			Email:             "user@example.com",
			Phone:             "+00 000 000000",
			Location:          "City, Country",
			SocialLink: &SocialLink{
				Linkedin: "https://linkedin.com/in/username",
				Github:   "https://github.com/username",
				Telegram: "https://t.me/username",
			},
		},
		Experiences: []Experience{
			{
				ID:          0,
				JobTitle:    "Senior Software Engineer",
				Company:     "Example Company",
				Location:    "City, Country",
				StartDate:   time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				Description: "Designing and maintaining backend services and distributed systems.",
			},
			{
				ID:          1,
				JobTitle:    "Software Engineer",
				Company:     "Another Company",
				Location:    "Remote",
				StartDate:   time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     util.Ptr(time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC)),
				Description: "Developed APIs, optimized database queries, and improved system reliability.",
			},
		},
		Skills: Skills{
			Technical: []string{
				"Programming Language",
				"Database",
				"Containerization",
				"Cloud Services",
				"API Development",
			},
			Soft: []string{
				"Problem Solving",
				"Communication",
				"Teamwork",
			},
		},
		Educations: []Education{
			{
				ID: 0,
				Degree:        "Bachelor's Degree in Computer Science",
				School:        "Example University",
				Location:      "City, Country",
				StartDate:     time.Date(2015, time.September, 1, 0, 0, 0, 0, time.UTC),
				Description:   "Studied software development, algorithms, and system design.",
			},
		},
		Projects: []Project{
			{
				Title:            "Sample Project",
				ShortDescription: "A scalable backend application",
				Description:      "Built a scalable system designed to handle high traffic and ensure reliability.\n\nIncludes monitoring, logging, and automated deployments.",
				Tags:             []string{"Backend", "Database", "Cloud"},
				LiveURL:          util.Ptr("https://example.com"),
				RepositoryURL:    util.Ptr("https://github.com/username/sample-project"),
				MyRole:           "Backend Developer",
				Duration:         "3 months",
				Client:           "Client Organization",
				Features: []*ProjectFeature{
					{
						ID: 0,
						Title:       "Scalable Architecture",
						Description: "Designed for high availability and horizontal scaling",
						Icon:        "analytics",
					},
					{
						ID: 1,
						Title:       "Resilient System",
						Description: "Implements retries and fault tolerance mechanisms",
						Icon:        "security",
					},
				},
				TechStack: map[string][]string{
					"Backend":  []string{"Language", "Framework"},
					"Database": []string{"SQL Database", "Cache"},
					"DevOps":   []string{"Containers", "Orchestration"},
				},
			},
		},
	}
}
