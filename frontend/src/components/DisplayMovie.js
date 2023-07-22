import Movie from "../Movie"

const DisplayMovie =(props)=>{
    let allmovies = props.movies
    return (
       <div className="allmovies">
          {
            allmovies.map(x =>{
                return <Movie movie={x}/>
            })
          }
       </div>
    )
} 
export default DisplayMovie;