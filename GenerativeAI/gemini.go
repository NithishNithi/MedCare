package generativeai

import (
	"context"
	"fmt"
	"log"
	"medcare/constants"
	"medcare/models"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const terminology = `
Anesthesiology
Cardiology
Dermatology
Emergency Medicine
Endocrinology
Family Medicine
Gastroenterology
Geriatrics
Hematology
Infectious Disease
Internal Medicine
Nephrology
Neurology
Obstetrics and Gynecology (OB/GYN)
Oncology
Ophthalmology
Orthopedics
Otolaryngology (ENT)
Pediatrics
Physical Medicine and Rehabilitation
Psychiatry
Pulmonology
Radiology
Rheumatology
Urology
Allergy and Immunology
Andrology
Bariatric Medicine
Cardiothoracic Surgery
Child and Adolescent Psychiatry
Clinical Genetics
Colorectal Surgery
Critical Care Medicine
Dermatopathology
Developmental-Behavioral Pediatrics
Diagnostic Radiology
Emergency Medical Services
Endocrine Surgery
Endovascular Surgical Neuroradiology
Facial Plastic Surgery
Forensic Pathology
Forensic Psychiatry
Gastrointestinal Surgery
General Surgery
Geriatric Psychiatry
Gynecologic Oncology
Hand Surgery
Hepatology
Hospice and Palliative Medicine
Hyperbaric Medicine
Infectious Disease Medicine
Interventional Cardiology
Interventional Radiology
Maternal-Fetal Medicine
Medical Biochemical Genetics
Medical Genetics
Medical Oncology
Medical Toxicology
Molecular Genetic Pathology
Neonatal-Perinatal Medicine
Neurological Surgery
Neuropathology
Neuroradiology
Nuclear Medicine
Nuclear Radiology
Obstetric Anesthesiology
Occupational Medicine
Ophthalmic Plastic and Reconstructive Surgery
Orthopedic Sports Medicine
Orthopedic Surgery
Orthopedic Trauma
Pain Medicine
Pediatric Anesthesiology
Pediatric Cardiology
Pediatric Critical Care Medicine
Pediatric Emergency Medicine
Pediatric Endocrinology
Pediatric Gastroenterology
Pediatric Hematology-Oncology
Pediatric Infectious Diseases
Pediatric Nephrology
Pediatric Pulmonology
Pediatric Radiology
Pediatric Rehabilitation Medicine
Pediatric Rheumatology
Pediatric Surgery
Pediatric Transplant Hepatology
Pediatric Urology
Plastic Surgery
Preventive Medicine
Psychiatry (General)
Psychosomatic Medicine
Public Health and General Preventive Medicine
Pulmonary Critical Care Medicine
Radiation Oncology
Reproductive Endocrinology and Infertility
Sleep Medicine
Spinal Cord Injury Medicine
Sports Medicine
Surgical Critical Care
Thoracic Surgery
Transplant Hepatology
Undersea and Hyperbaric Medicine
Urgent Care Medicine
Vascular Neurology
Vascular Surgery
Addiction Medicine
Aerospace Medicine
Clinical Cardiac Electrophysiology
Clinical Informatics
Correctional Medicine
Craniofacial Surgery
Cytopathology
Disaster Medicine
Epilepsy
Female Pelvic Medicine and Reconstructive Surgery
Forensic Medicine
Hospital Medicine
Infectious Diseases
Integrative Medicine
Medical Microbiology
Medical Physics
Neuromuscular Medicine
Pediatric Pathology
Selective Pathology
Surgical Oncology
Transfusion Medicine
Trauma Surgery
Traumatology
Tropical Medicine
Vascular Medicine
Venous and Lymphatic Medicine`

func GenerativeAI(request *models.BookAppointment) (string, error) {
	ctx := context.Background()
	// Set up the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(constants.Gemini_apiKey))
	if err != nil {
		log.Println(err)
		return "", err

	}
	defer client.Close()
	// Specify the model to use
	model := client.GenerativeModel("gemini-pro")
	// Generate content
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("I have symptoms: %s. Description: %s. Please find suitable terminologies from medical terminologies: %s", request.Symptoms, request.BriefDescription, terminology+"not more than one terminolgy"+"Note:Only give the terminology")))
	if err != nil {
		log.Println(err)
		return "", err
	}
	var result string
	for _, item := range resp.Candidates {
		fmt.Println("Generated content:", item.Content.Parts[0])
		result = (RetrieveResponce(item.Content.Parts))

	}
	return result, nil
}

func RetrieveResponce(parts []genai.Part) string {
	var result string
	for _, part := range parts {
		result += fmt.Sprint(part)
	}
	return result
}
