import React, { useEffect } from "react";
import clsx from "clsx";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import AppBar from "@material-ui/core/AppBar";
import Tooltip from "@material-ui/core/Tooltip";
import Toolbar from "@material-ui/core/Toolbar";
import List from "@material-ui/core/List";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import Button from "@material-ui/core/Button";

import HomeIcon from "@material-ui/icons/Home";
import AccountCircleIcon from "@material-ui/icons/AccountCircle";
import AddCircleOutlineIcon from "@material-ui/icons/AddCircleOutline";
import AssignmentReturnIcon from '@material-ui/icons/AssignmentReturn';

//import Navbar from "./components/Navbar";

//import UserCreate from "./components/UserCreate";
import BookInformations from "./components/BookInformation";
import BookInformationCreate from "./components/BookInformationCreate";
import SignIn from "./components/SignIn";
import Home from "./components/Home";
import BookOrders from "./components/BookOrders";
import BookOrderCreate from "./components/BookOrderCreate";
import ShoppingCartOutlinedIcon from '@material-ui/icons/ShoppingCartOutlined';
import BookIcon from '@material-ui/icons/Book';
import TabletAndroidIcon from '@material-ui/icons/TabletAndroid';
import DeviceBorrows from "./components/DeviceBorrows";
import DeviceBorrowCreate from "./components/DeviceBorrowCreate";
import BookReturnCreate from "./components/BookReturnCreate";
import BookReturns from "./components/BookReturns";
import CreateBorrowDetail from "./components/BorrowDetailCreate";
import Borrows from "./components/BorrowDetails";
import BookingRoom from "./components/BookingRoom";
import BookingRoomCreate from "./components/BookingRoomCreate";
import Research from "./components/Research";
import ResearchCreate from "./components/ResearchCreate";
import AccountCircle from '@material-ui/icons/AccountCircle';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import { MembersInterface } from "./models/IMember";
import Chip from '@material-ui/core/Chip';

const drawerWidth = 240;
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },
    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3)
    },
    a: {
      textDecoration: "none",
      color: "inherit",
    },
  })
);
export default function MiniDrawer() {
  const classes = useStyles();
  const theme = useTheme();
  const [member, setMember] = React.useState<MembersInterface>();
  const Role = localStorage.getItem("role");
  const [open, setOpen] = React.useState(false);
  const [token, setToken] = React.useState<String>("");
  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { Authorization: `Bearer ${localStorage.getItem("token")}`, "Content-Type": "application/json" },
  };

  const getMember = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/member/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log("member", res.data);
        if (res.data) {
          setMember(res.data);
        } else {
          console.log("else");
        }
      });
  };


  const menuLibrarian = [
    { name: "?????????????????????", icon: <HomeIcon />, path: "/" },
    { name: "???????????????????????????????????????????????????????????????", icon: <ShoppingCartOutlinedIcon />, path: "/bookorders" },
    { name: "????????????????????????????????????????????????", icon: <AddCircleOutlineIcon />, path: "/book_informations" },
    { name: "??????????????????????????????????????????", icon: <AddCircleOutlineIcon />, path: "/Research" },
  ];

  const menuMember = [
    { name: "?????????????????????", icon: <HomeIcon />, path: "/" },
    { name: "?????????????????????????????????????????????????????????", icon: <TabletAndroidIcon />, path: "/deviceborrows" },
    { name: "?????????????????????????????????????????????????????????", icon: <AccountCircleIcon />, path: "/booking_rooms" },
    { name: "?????????????????????????????????????????????????????????", icon: <BookIcon />, path: "/borrowDetail" },
    { name: "?????????????????????????????????????????????????????????", icon: <AssignmentReturnIcon />, path: "/book_return" },
  ];

  const Menu2 = () => {
    if (Role === "??????????????????????????????") {
      return menuLibrarian
    } else {
      return menuMember
    }
  }

  useEffect(() => {
    const token = localStorage.getItem("token");
    getMember();
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };


  return (
    <div className={classes.root}>
      <Router>
        <CssBaseline />
        {token && (
          <>
            <AppBar
              position="fixed"
              style={{ backgroundColor: '#13c2c2' }}
              className={clsx(classes.appBar, {
                [classes.appBarShift]: open,
              })}
            >
              <Toolbar>
                <IconButton
                  color="inherit"
                  aria-label="open drawer"
                  onClick={handleDrawerOpen}
                  edge="start"
                  className={clsx(classes.menuButton, {
                    [classes.hide]: open,
                  })}
                >
                  <MenuIcon />
                </IconButton>
                <Typography variant="h6" className={classes.title}>
                  ????????????????????????????????????
                </Typography>
                
                <Chip
                  size="medium"
                  icon={<AccountCircleIcon style={{ color: '#13c2c2' }} />}
                  label={member?.Name + " (" + member?.Role + ") "}
                  variant="outlined"
                  style={{ backgroundColor: '#fff', fontSize: '1rem', color: '#009688' }}
                />
                <Tooltip title = "Logout">
                <Button color="inherit" onClick={signout} style={{ fontFamily: "Kanit" }}>
                  <ExitToAppIcon style={{ fontSize: 30, marginRight: 2 }} />
                </Button>
                </Tooltip>

              </Toolbar>
            </AppBar>
            <Drawer
              variant="permanent"
              className={clsx(classes.drawer, {
                [classes.drawerOpen]: open,
                [classes.drawerClose]: !open,
              })}
              classes={{
                paper: clsx({
                  [classes.drawerOpen]: open,
                  [classes.drawerClose]: !open,
                }),
              }}
            >
              <div className={classes.toolbar}>
                <IconButton onClick={handleDrawerClose}>
                  {theme.direction === "rtl" ? (
                    <ChevronRightIcon />
                  ) : (
                    <ChevronLeftIcon />
                  )}
                </IconButton>
              </div>
              <Divider />
              <List>
                {Menu2().map((item, index) => (
                  <Link to={item.path} key={item.name} className={classes.a}>
                    <ListItem button>
                      <ListItemIcon>{item.icon}</ListItemIcon>
                      <ListItemText primary={item.name} />
                    </ListItem>
                  </Link>
                ))}
              </List>
            </Drawer>
          </>
        )}
        <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/borrowDetail/create" component={CreateBorrowDetail} />
              <Route exact path="/borrowDetail" component={Borrows} />
              <Route exact path="/book_return" component={BookReturns} />
              <Route exact path="/book_return/create" component={BookReturnCreate} />
              <Route exact path="/book_informations/create" component={BookInformationCreate} />
              <Route exact path="/book_informations" component={BookInformations} />
              <Route exact path="/bookorders" component={BookOrders} />
              <Route exact path="/bookordercreate" component={BookOrderCreate} />
              <Route exact path="/deviceborrows" component={DeviceBorrows} />
              <Route exact path="/deviceborrow/create" component={DeviceBorrowCreate} />
              <Route exact path="/booking_roomscreate" component={BookingRoomCreate} />
              <Route exact path="/booking_rooms" component={BookingRoom} />
              <Route exact path="/researCreate/create" component={ResearchCreate} />
              <Route exact path="/Research" component={Research} />
            </Switch>
          </div>
        </main>
      </Router>
    </div>
  );
}