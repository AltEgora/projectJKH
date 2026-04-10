namespace GkhCalculator.Models
{
    public class GkhTariffsOptions
    {
        public const string SectionName = "GkhTariffs";

        public decimal ColdWaterTariff { get; set; }
        /// <summary>ГВС: компонент за 1 м³ (холодная составляющая / поставка).</summary>
        public decimal HotWaterVolumeTariff { get; set; }
        /// <summary>ГВС: подогрев за 1 Гкал.</summary>
        public decimal HotWaterHeatTariff { get; set; }
        public decimal SewerTariff { get; set; }
        public decimal ElectricityTariff { get; set; }
        public decimal HeatingTariff { get; set; }
        public decimal GasTariff { get; set; }
        public decimal WasteTariffPerPerson { get; set; }

        public decimal ColdWaterNormM3PerPerson { get; set; }

        public decimal HotWaterNormM3PerPerson { get; set; }
        public decimal GasNormM3PerPerson { get; set; }
        public decimal HeatingNormGcalPerM2 { get; set; }

        public decimal ElectricityNightTariff { get; set; }
        public decimal CapitalRepairTariffPerM2 { get; set; }
        public decimal OdnFraction { get; set; }
        public decimal CommercialResourceMultiplier { get; set; }
    }
}
