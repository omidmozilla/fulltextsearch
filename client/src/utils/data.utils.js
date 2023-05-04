export function IsEmptyString(value) {
  if( value.trim().length == 0 || value.trim() == "" ) {
    return true 
  }
  
  return false
}
