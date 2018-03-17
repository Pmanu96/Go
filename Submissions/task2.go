+import "fmt"
+
+  func sum(i int)
+  {
+      add = +i
+  }
+
+func main()
+{
+    for i:=1; i<=5; i++
+    {
+      fmt.Print("  ",i)
+       go sum(i)
+        sum(i)
+    }
+    fmt.Print(add)
+}
