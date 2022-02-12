import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import { ImageList, styled, Paper, Grid, Box } from "@material-ui/core";
import ImageListItem from '@material-ui/core/ImageListItem';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);
const Item = styled(Paper)(({ theme }) => ({
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));
function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบห้องสมุด</h1>
        <img src="https://images.pexels.com/photos/1319855/pexels-photo-1319855.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940"
          width="900px"></img>
      </Container>
    </div>
  );
}
export default Home;
