# acb_task

A simple Movie APIs with GET and POST method. 
The Rest APIs could be accessed via: https://acb-movie-api-56eyiclpoa-uc.a.run.app (powered by Google Cloud Run)

List of APIs:
1. GetMovies [GET] - (/get_movies): Get all movies data
2. GetMovie [GET] - (/get_movie/{id}): Get specific movie data
3. AddMovie [POST] - (/add_movie): Adding new movie 
   Request Body Example (json):
   {
        "name": "Up",
        "genre": "Animation",
        "release_year": 2012,
        "production_house": "Disney"
   }
4. HealthCheckAPI [GET] - (/api/health): Health Check API
