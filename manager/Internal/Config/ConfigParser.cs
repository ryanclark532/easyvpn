
namespace manager.Internal.Config;

public class ConfigParser(string configFile)
{
    static readonly HashSet<string> PrivateConfigItems = ["ca", "key", "cert", "dh", "auth-user-pass-verify"];

    private readonly string _configFile = configFile ?? throw new ArgumentNullException(nameof(configFile));
    private readonly Dictionary<string, string> _config = new();

    public Dictionary<string, string> Config =>
        _config.Where(kvp => !PrivateConfigItems.Contains(kvp.Key))
            .ToDictionary(kvp => kvp.Key, kvp => kvp.Value);

    public void ReadConfig()
    {
        if (string.IsNullOrEmpty(_configFile))
        {
            throw new FileNotFoundException("Config path not set");
        }
        if (!File.Exists(_configFile))
        {
            throw new FileNotFoundException("The config file could not be found.");
        }
        foreach (var line in File.ReadLines(_configFile))
        {
            string trimmedLine = line.Trim();

            if (string.IsNullOrEmpty(trimmedLine) || trimmedLine.StartsWith("#"))
                continue;

            var parts = trimmedLine.Split([' '], 2, StringSplitOptions.RemoveEmptyEntries);

            if (parts.Length != 2)
            {
                continue;
            }

            var key = parts[0];
            var value = parts[1];
            _config[key] = value;
        }
    }

    public void WriteConfig()
    {
        
    }
}