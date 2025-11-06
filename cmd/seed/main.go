package main

import (
	"log"
	"sportTrackerAPI/db"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/internal/exercise"
)

func main() {
	cfg := config.LoadConfig()
	database := db.NewDb(cfg)

	exercises := []exercise.Exercise{
		// CHEST
		{
			Name:        "Жим штанги лежа",
			MuscleGroup: "Chest",
			Description: "Базовое упражнение для развития грудных мышц",
			Instruction: "Лежа на горизонтальной скамье, опустите штангу к груди и мощно выжмите вверх. Локти под углом 45 градусов.",
			Difficulty:  3,
		},
		{
			Name:        "Жим гантелей на наклонной скамье",
			MuscleGroup: "Chest",
			Description: "Упражнение для верхней части грудных мышц",
			Instruction: "На наклонной скамье (30-45 градусов) жмите гантели вверх, сводя их в верхней точке.",
			Difficulty:  2,
		},
		{
			Name:        "Отжимания на брусьях",
			MuscleGroup: "Chest",
			Description: "Упражнение с весом тела для грудных и трицепсов",
			Instruction: "На брусьях опуститесь вниз, наклонив корпус вперед, затем мощно поднимитесь.",
			Difficulty:  3,
		},
		{
			Name:        "Сведения в кроссовере",
			MuscleGroup: "Chest",
			Description: "Изолирующее упражнение для грудных мышц",
			Instruction: "Стоя между стойками кроссовера, сведите рукоятки перед собой по дугообразной траектории.",
			Difficulty:  2,
		},

		// BACK
		{
			Name:        "Становая тяга",
			MuscleGroup: "Back",
			Description: "Фундаментальное упражнение для всей задней цепи",
			Instruction: "Ноги на ширине плеч, возьмите штангу прямым хватом. С прямой спиной поднимите штангу за счет усилия ног и спины.",
			Difficulty:  5,
		},
		{
			Name:        "Подтягивания широким хватом",
			MuscleGroup: "Back",
			Description: "Базовое упражнение для широчайших мышц",
			Instruction: "Широким хватом повисните на перекладине, подтянитесь до уровня подбородка, чувствуя работу широчайших.",
			Difficulty:  4,
		},
		{
			Name:        "Тяга штанги в наклоне",
			MuscleGroup: "Back",
			Description: "Упражнение для толщины спины",
			Instruction: "В наклоне с прямой спиной подтяните штангу к поясу, сводя лопатки.",
			Difficulty:  3,
		},
		{
			Name:        "Тяга верхнего блока",
			MuscleGroup: "Back",
			Description: "Аналог подтягиваний для развития широчайших",
			Instruction: "Сидя у тренажера, тяните рукоять к груди, отклоняя корпус немного назад.",
			Difficulty:  2,
		},

		// BICEPS
		{
			Name:        "Подъем штанги на бицепс",
			MuscleGroup: "Biceps",
			Description: "Базовое упражнение для бицепса",
			Instruction: "Стоя, поднимите штангу на бицепс, не раскачивая корпус. Локти прижаты к туловищу.",
			Difficulty:  2,
		},
		{
			Name:        "Подъем гантелей на бицепс сидя",
			MuscleGroup: "Biceps",
			Description: "Изолирующая работа на бицепс",
			Instruction: "Сидя на скамье, попеременно поднимайте гантели на бицепс с супинацией кисти.",
			Difficulty:  2,
		},
		{
			Name:        "Молотки",
			MuscleGroup: "Biceps",
			Description: "Упражнение для брахиалиса и бицепса",
			Instruction: "Поднимайте гантели нейтральным хватом (ладони смотрят друг на друга).",
			Difficulty:  2,
		},
		{
			Name:        "Концентрированные сгибания",
			MuscleGroup: "Biceps",
			Description: "Изолирующее упражнение на пик бицепса",
			Instruction: "Сидя, уприте локоть во внутреннюю часть бедра и сгибайте руку с гантелью.",
			Difficulty:  1,
		},

		// TRICEPS
		{
			Name:        "Жим лежа узким хватом",
			MuscleGroup: "Triceps",
			Description: "Базовое упражнение для трицепса",
			Instruction: "Лежа на скамье, жмите штангу узким хватом, локти прижаты к туловищу.",
			Difficulty:  3,
		},
		{
			Name:        "Французский жим",
			MuscleGroup: "Triceps",
			Description: "Изолирующее упражнение для трицепса",
			Instruction: "Лежа на скамье, опускайте штангу ко лбу, сгибая руки только в локтевых суставах.",
			Difficulty:  2,
		},
		{
			Name:        "Разгибания на блоке",
			MuscleGroup: "Triceps",
			Description: "Упражнение для проработки всех головок трицепса",
			Instruction: "Стоя у блока, разгибайте руки вниз, держа локти неподвижно.",
			Difficulty:  1,
		},
		{
			Name:        "Отжимания от скамьи",
			MuscleGroup: "Triceps",
			Description: "Упражнение с весом тела для трицепса",
			Instruction: "Спиной к скамье, опускайтесь вниз, сгибая руки в локтях до прямого угла.",
			Difficulty:  2,
		},

		// SHOULDERS
		{
			Name:        "Армейский жим",
			MuscleGroup: "Shoulders",
			Description: "Базовое упражнение для дельтовидных мышц",
			Instruction: "Стоя или сидя, жмите штангу с груди вверх над головой.",
			Difficulty:  3,
		},
		{
			Name:        "Жим гантелей сидя",
			MuscleGroup: "Shoulders",
			Description: "Альтернатива армейскому жиму",
			Instruction: "Сидя на скамье, жмите гантели вверх по широкой дуге.",
			Difficulty:  2,
		},
		{
			Name:        "Махи гантелями в стороны",
			MuscleGroup: "Shoulders",
			Description: "Для средних пучков дельт",
			Instruction: "Стоя, поднимайте гантели через стороны до уровня плеч.",
			Difficulty:  2,
		},
		{
			Name:        "Подъем гантелей перед собой",
			MuscleGroup: "Shoulders",
			Description: "Для передних пучков дельт",
			Instruction: "Поднимайте гантели перед собой попеременно или вместе.",
			Difficulty:  1,
		},

		// LEGS
		{
			Name:        "Приседания со штангой",
			MuscleGroup: "Legs",
			Description: "Король упражнений для ног",
			Instruction: "Штанга на трапециях, приседайте до параллели бедер с полом, колени не выходят за носки.",
			Difficulty:  4,
		},
		{
			Name:        "Жим ногами",
			MuscleGroup: "Legs",
			Description: "Упражнение для квадрицепсов и ягодиц",
			Instruction: "В тренажере жмите платформу, не разгибая колени полностью в верхней точке.",
			Difficulty:  2,
		},
		{
			Name:        "Выпады с гантелями",
			MuscleGroup: "Legs",
			Description: "Упражнение для ног и ягодиц",
			Instruction: "С гантелями в руках делайте шаг вперед и опускайтесь до прямого угла в коленях.",
			Difficulty:  3,
		},
		{
			Name:        "Румынская тяга",
			MuscleGroup: "Legs",
			Description: "Для бицепса бедра и ягодиц",
			Instruction: "На слегка согнутых ногах наклонитесь со штангой, чувствуя растяжение бицепса бедра.",
			Difficulty:  3,
		},
		{
			Name:        "Подъем на носки стоя",
			MuscleGroup: "Legs",
			Description: "Для икроножных мышц",
			Instruction: "В тренажере поднимайтесь на носки, максимально растягивая икры.",
			Difficulty:  1,
		},
	}

	result := database.Create(exercises)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
