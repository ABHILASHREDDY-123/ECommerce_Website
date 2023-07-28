const SearchByName=(name)=>{
    
}
const Search = () => {
  return (
    <div>
      <input type="text" placeholder="Search.." onChange={(e)=>{SearchByName(e.target.value)}}/>
    </div>
  );
};
export default Search;