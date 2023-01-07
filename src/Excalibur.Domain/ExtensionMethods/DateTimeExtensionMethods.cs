namespace Excalibur.Domain.ExtensionMethods;

using System.Globalization;

public static class DateTimeExtensionMethods
{
    /// <summary>
    /// Returns a string representation of the DateTime object in the ISO 8601 format.
    /// </summary>
    public static string ToIsoFormatString(this DateTime dateTime)
    {
        return dateTime.ToString("yyyy-MM-ddTHH\\:mm\\:ss", CultureInfo.InvariantCulture);
    }
}
