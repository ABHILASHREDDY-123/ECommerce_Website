import { useState,useEffect } from "react"
import axios from "axios"
const Expand =(id,setFullmovie)=>{
    const link=`https://www.omdbapi.com/?i=${id}&apikey=654293a2`;
    console.log(link);
    axios.get(link).then((res)=>{
      console.log(res);
       if(res.data.Title)setFullmovie(res.data);
    })
}
const Movie = (props)=>{
   let each_movie=props.movie
   const [full_movie,setFullmovie]=useState({})
   useEffect(() => {
       setFullmovie({});
    }, [props.movie]);
   return (
     <div className="movie">
        <img src={each_movie.Poster} height="300px" onClick={()=>{Expand(each_movie.imdbID,setFullmovie)}}/>
        <h3>{each_movie.Title}</h3>
        <h3>{each_movie.Year}</h3>
        {
         full_movie.Title?(
            <div>
               <h3>{full_movie.Released}</h3>
               <h3>{full_movie.imdbRating}</h3>
               <h3>{full_movie.Language}</h3>
            </div>
         ):<></>
        }
     </div>
   ) 
}
export default Movie