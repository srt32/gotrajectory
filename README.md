gotrajectory
============

A WIP golang wrapper for [apptrajectory](https://www.apptrajectory.com/)


Example response:

```
{
stories:
  [
    {
      archived: false,
      assignee_id: null,
      branch: null,
      comments_count: 1,
      created_at: "2014-05-30T17:15:10-07:00",
      deleted: false,
      design_needed: true,
      development_needed: true,
      id: 15644034,
      idea_id: null,
      iteration_id: 16515878,
      points: 1,
      position: 100,
      state: "unstarted",
      task_type: "Feature",
      title: "Whoohoo! My first trajectory",
      updated_at: "2014-05-30T17:15:10-07:00",
      user_id: 15508297,
      state_events:
        [
          "start"
        ],
      comma_separated_tags: null
  }
  ],
    iterations:
      [
        {
          accepted_points: 0,
          comments_count: 0,
          complete: false,
          created_at: "2014-05-30T17:07:23-07:00",
          estimated_points: 1,
          estimated_velocity: 20,
          id: 16515878,
          starts_on: "2014-05-26",
          stories_count: 1,
          team_strength: 100,
          updated_at: "2014-05-30T17:15:10-07:00",
          percent_complete: 0,
          current?: true,
          future?: false,
          accepted_stories_count: 0
      }
    ]
}
```
