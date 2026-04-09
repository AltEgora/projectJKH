namespace GkhCalculator.Models
{
    /// <summary>Профиль из опроса (index.html), влияет на способ расчёта.</summary>
    public class GkhProfile
    {
        public string PropertyType { get; set; } = "residential";
        public string BuildingType { get; set; } = "apartment";
        public string HeatingType { get; set; } = "central";
        public bool HeatingMeter { get; set; }
        public bool WaterMeter { get; set; }
        public string ElectricityType { get; set; } = "single";
        /// <summary>gas — базовый тариф; electric — дневной и ночной тарифы × 0,7.</summary>
        public string StoveType { get; set; } = "gas";
        public string GasType { get; set; } = "none";
        public bool Odn { get; set; }
        public decimal Area { get; set; }
        public int Residents { get; set; }
    }
}
