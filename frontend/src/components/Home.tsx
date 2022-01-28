import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

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

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบบันทึกข้อมูลหนังสือ</h1>
        <h3>Requirements</h3>
        <p>
        ระบบห้องสมุด เป็นระบบที่ให้ผู้ใช้ระบบซึ่งเข้าระบบในฐานะของสมาชิกของห้องสมุด จะสามารถเข้าระบบเพื่อใช้บริการต่างๆของห้องสมุดได้ โดยสมาชิกแต่ละคนสามารถที่จะยืมหนังสือ 
        คืนหนังสือ และยังสามารถจองห้องทบทวนได้อีกด้วย และในส่วนของผู้ใช้ระบบที่เข้าระบบในฐานะบรรณารักษ์ห้องสมุดจะสามารถสั่งซื้อหนังสือ บันทึกข้อมูลงานวิจัย และยังสามารถยืมอุปกรณ์และคืนอุปกรณ์ได้
        และนอกจากความสามารถเหล่านี้แล้วยังมีระบบที่สามารถบันทึกข้อมูลของหนังสือ เป็นระบบที่เป็นการจัดประเภทหนังสือ เพื่อให้เกิดความสะดวกแก่ผู้มาใช้บริการห้องสมุด โดยการบันทึกข้อมูลของหนังสือ 
        จะสามารถบันทึกประเภทของหนังสือ สถานที่จัดเก็บ ข้อมูลของหนังสือ เช่น เลขเรียกหนังสือ ชื่อเรื่อง ผู้แต่ง ปีที่พิมพ์ และยังสามารถบันทึกวันที่และเวลาที่ทำการบันทึกได้อีกด้วย

        </p>
      </Container>
    </div>
  );
}
export default Home;