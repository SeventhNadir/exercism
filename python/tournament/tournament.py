class Team:
    def __init__(self, name):
        self.name = name
        self.won = 0
        self.tied = 0
        self.lost = 0

    def played(self):
        return self.won + self.tied + self.lost

    def points(self):
        return 3 * self.won + self.tied

    def __lt__(self, other):
        if self.points() == other.points():
            return self.name > other.name
        else:
            return self.points() < other.points()


class Leaderboard:
    def __init__(self):
        self.teams = {}

    def get_team(self, team):
        if self.teams.get(team.name):
            return self.teams.get(team.name)
        else:
            self.teams[team.name] = team
            return self.teams.get(team.name)

    def table(self):
        results = [f"{'Team': <31}| MP |  W |  D |  L |  P"]
        for team in sorted(self.teams.values(), reverse=True):
            results.append(
                f"{team.name: <31}|  {team.played()} |  {team.won} |  {team.tied} |  {team.lost} |  {team.points()}"
            )
        return results


def tally(rows):
    leaderboard = Leaderboard()

    for row in rows:
        teamA, teamB, result = row.split(";")
        teamA, teamB = Team(teamA), Team(teamB)
        teamA, teamB = leaderboard.get_team(teamA), leaderboard.get_team(teamB)
        if result == "win":
            teamA.won += 1
            teamB.lost +=1
        if result == "draw":
            teamA.tied += 1
            teamB.tied += 1
        if result == "loss":
            teamA.lost += 1
            teamB.won += 1
    return leaderboard.table()
