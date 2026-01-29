package cn.xswitch.proto;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import cman.CManProto;
import org.junit.jupiter.api.Test;
import xctrl.XCtrl;

public class ProtoSmokeTest {
  @Test
  void descriptorsAreLoadable() {
    assertNotNull(CManProto.getDescriptor());
    assertNotNull(XCtrl.getDescriptor());
    assertEquals("cman", CManProto.getDescriptor().getPackage());
    assertEquals("xctrl", XCtrl.getDescriptor().getPackage());
  }
}
