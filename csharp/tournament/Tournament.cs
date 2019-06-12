using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Collections.Generic;
public static class Tournament
{   
    public static void Tally(Stream inStream, Stream outStream)
    {
        var teams = new Dictionary<string, Team>();

        StreamReader reader = new StreamReader(inStream);
        string lines = reader.ReadToEnd();
        string[] records = lines.Split("\n");
        reader.Dispose();
        inStream.Dispose();

        foreach (string record in records)
        {
            string[] match = record.Split(";");
            if (record == "") 
            {
                break;
            }
            string outcome = match[2];
            string firstTeam = match[0];
            string secondTeam = match[1];

            var team1 = GetOrCreateTeam(firstTeam, teams);
            var team2 = GetOrCreateTeam(secondTeam, teams);

            team1.played++;
            team2.played++;
            switch (outcome)
            {
                case "win":
                    team1.won++;
                    team2.lost++;
                    break;
                case "loss":
                    team1.lost++;
                    team2.won++;
                    break;
                case "draw":
                    team1.drawn++;
                    team2.drawn++;
                    break;
            }

            teams[firstTeam] = team1;
            teams[secondTeam] = team2;
        }

        var header = "Team                           | MP |  W |  D |  L |  P";
        var template = "{0,-30} | {1,2} | {2,2} | {3,2} | {4,2} | {5,2}";

        var sorted = teams.Values
            .OrderByDescending(t => t.won*3 + t.drawn)
            .ThenBy(t => t.name)
            .Select(t => String.Format(template, t.name, t.played, t.won, t.drawn, t.lost, (t.won*3 + t.drawn)));
        sorted = sorted.Prepend(header);

        outStream.Write(Encoding.UTF8.GetBytes(String.Join("\n", sorted)));
    }
    
    static Team GetOrCreateTeam(string name, Dictionary<string, Team> teams)
    {
        try
        {
            return teams[name];
        }
        catch (KeyNotFoundException)
        {
            return new Team(name);
        }
    }

    public class Team
    {
        public string name;
        public int played = 0;
        public int won = 0;
        public int drawn = 0;
        public int lost = 0;

        public Team(string Name)
        {
            name = Name;
        }
    }
}