namespace GkhCalculator.Models
{
    public class GkhRequest
    {
        public GkhProfile? Profile { get; set; }

        public decimal ColdWaterVolume { get; set; }
        public decimal HotWaterVolume { get; set; }
        /// <summary>Гкал на подогрев ГВС (отдельное показание в квитанции).</summary>
        public decimal HotWaterThermalGcal { get; set; }
        public decimal ElectricityVolume { get; set; }
        public decimal ElectricityDayVolume { get; set; }
        public decimal ElectricityNightVolume { get; set; }
        public decimal HeatingVolume { get; set; }
        public decimal GasVolume { get; set; }

        public int Residents { get; set; }

        public bool HasBenefits { get; set; }

        public decimal UserTotal { get; set; }

        /// <summary>
        /// Для узких калькуляторов: cold | hot | heat — сравнить сумму строки квитанции с подытогом по этой услуге.
        /// </summary>
        public string? ReceiptLineMode { get; set; }

        /// <summary>Сумма по строке квитанции для сравнения (при заданном ReceiptLineMode).</summary>
        public decimal ReceiptLineAmount { get; set; }
    }
}