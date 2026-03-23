using GkhCalculator.Models;

namespace GkhCalculator.Services
{
    public class GkhCalculatorService
    {
        private const decimal ColdWaterTariff = 30m;
        private const decimal SewerTariff = 25m;
        private const decimal ElectricityTariff = 5m;
        private const decimal HeatingTariff = 2200m;

        public object Calculate(GkhRequest req)
        {
            var coldWater = req.ColdWaterVolume * ColdWaterTariff;
            var sewer = req.SewerVolume * SewerTariff;
            var electricity = req.ElectricityVolume * ElectricityTariff;
            var heating = req.HeatingVolume * HeatingTariff;

            var total = coldWater + sewer + electricity + heating;

            if (req.HasBenefits)
                total *= 0.7m;

            var diff = req.UserTotal - total;

            string status = Math.Abs(diff) < 1
                ? "ok"
                : diff > 0 ? "overpay" : "underpay";

            return new
            {
                total,
                difference = diff,
                status,
                breakdown = new
                {
                    coldWater,
                    sewer,
                    electricity,
                    heating
                }
            };
        }
    }
}