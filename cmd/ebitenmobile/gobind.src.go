// Code generated by file2byteslice. DO NOT EDIT.
// (gofmt is fine after generating)

package main

var gobindsrc = []byte("// Copyright 2019 The Ebiten Authors\n//\n// Licensed under the Apache License, Version 2.0 (the \"License\");\n// you may not use this file except in compliance with the License.\n// You may obtain a copy of the License at\n//\n//     http://www.apache.org/licenses/LICENSE-2.0\n//\n// Unless required by applicable law or agreed to in writing, software\n// distributed under the License is distributed on an \"AS IS\" BASIS,\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n// See the License for the specific language governing permissions and\n// limitations under the License.\n\n// +build ebitenmobilegobind\n\n// gobind is a wrapper of the original gobind. This command adds extra files like a view controller.\npackage main\n\nimport (\n\t\"flag\"\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"log\"\n\t\"os\"\n\t\"os/exec\"\n\t\"path/filepath\"\n\t\"strings\"\n\n\t\"golang.org/x/tools/go/packages\"\n)\n\nvar (\n\tlang          = flag.String(\"lang\", \"\", \"\")\n\toutdir        = flag.String(\"outdir\", \"\", \"\")\n\tjavaPkg       = flag.String(\"javapkg\", \"\", \"\")\n\tprefix        = flag.String(\"prefix\", \"\", \"\")\n\tbootclasspath = flag.String(\"bootclasspath\", \"\", \"\")\n\tclasspath     = flag.String(\"classpath\", \"\", \"\")\n\ttags          = flag.String(\"tags\", \"\", \"\")\n)\n\nvar usage = `The Gobind tool generates Java language bindings for Go.\n\nFor usage details, see doc.go.`\n\nfunc main() {\n\tflag.Parse()\n\tif err := run(); err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n\nfunc run() error {\n\tcmd := exec.Command(\"gobind-original\", os.Args[1:]...)\n\tcmd.Stdout = os.Stdout\n\tcmd.Stderr = os.Stderr\n\tif err := cmd.Run(); err != nil {\n\t\treturn err\n\t}\n\n\tpkgs, err := packages.Load(nil, flag.Args()[0])\n\tif err != nil {\n\t\treturn err\n\t}\n\tprefixLower := *prefix + pkgs[0].Name\n\tprefixUpper := strings.Title(*prefix) + strings.Title(pkgs[0].Name)\n\n\twriteFile := func(filename string, content string) error {\n\t\tif err := ioutil.WriteFile(filepath.Join(*outdir, \"src\", \"gobind\", filename), []byte(content), 0644); err != nil {\n\t\t\treturn err\n\t\t}\n\t\treturn nil\n\t}\n\treplacePrefixes := func(content string) string {\n\t\tcontent = strings.ReplaceAll(content, \"{{.PrefixUpper}}\", prefixUpper)\n\t\tcontent = strings.ReplaceAll(content, \"{{.PrefixLower}}\", prefixLower)\n\t\treturn content\n\t}\n\n\t// Add additional files.\n\tlangs := strings.Split(*lang, \",\")\n\tfor _, lang := range langs {\n\t\tswitch lang {\n\t\tcase \"objc\":\n\t\t\t// iOS\n\t\t\tif err := writeFile(prefixLower+\"ebitenviewcontroller_ios.m\", replacePrefixes(objcM)); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\tcase \"java\":\n\t\t\t// Android\n\t\t\t// TODO: Insert a Java file and let the original gobind compile it.\n\t\tcase \"go\":\n\t\t\t// Do nothing.\n\t\tdefault:\n\t\t\tpanic(fmt.Sprintf(\"unsupported language: %s\", lang))\n\t\t}\n\t}\n\n\treturn nil\n}\n\nconst objcM = `// Code generated by ebitenmobile. DO NOT EDIT.\n\n// +build ios\n\n#import <stdint.h>\n#import <UIKit/UIKit.h>\n#import <GLKit/GLkit.h>\n#import \"Ebitenmobileview.objc.h\"\n\n@interface {{.PrefixUpper}}EbitenViewController : UIViewController\n@end\n\n@implementation {{.PrefixUpper}}EbitenViewController {\n  GLKView* glkView_;\n}\n\n- (GLKView*)glkView {\n  if (!glkView_) {\n    glkView_ = [[GLKView alloc] init];\n    glkView_.multipleTouchEnabled = YES;\n  }\n  return glkView_;\n}\n\n- (void)viewDidLoad {\n  [super viewDidLoad];\n\n  self.glkView.delegate = (id<GLKViewDelegate>)(self);\n  [self.view addSubview: self.glkView];\n\n  EAGLContext *context = [[EAGLContext alloc] initWithAPI:kEAGLRenderingAPIOpenGLES2];\n  [self glkView].context = context;\n\t\n  [EAGLContext setCurrentContext:context];\n\t\n  CADisplayLink *displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(drawFrame)];\n  [displayLink addToRunLoop:[NSRunLoop currentRunLoop] forMode:NSDefaultRunLoopMode];\n}\n\n- (void)viewDidLayoutSubviews {\n  [super viewDidLayoutSubviews];\n  CGRect viewRect = [[self view] frame];\n\n  EbitenmobileviewLayout(viewRect.size.width, viewRect.size.height, (id<EbitenmobileviewViewRectSetter>)self);\n}\n\n- (void)setViewRect:(long)x y:(long)y width:(long)width height:(long)height {\n  CGRect glkViewRect = CGRectMake(x, y, width, height);\n  [[self glkView] setFrame:glkViewRect];\n}\n\n- (void)didReceiveMemoryWarning {\n  [super didReceiveMemoryWarning];\n  // Dispose of any resources that can be recreated.\n  // TODO: Notify this to Go world?\n}\n\n- (void)drawFrame{\n  [[self glkView] setNeedsDisplay];\n}\n\n- (void)glkView:(GLKView*)view drawInRect:(CGRect)rect {\n  NSError* err = nil;\n  EbitenmobileviewUpdate(&err);\n  if (err != nil) {\n    NSLog(@\"Error: %@\", err);\n  }\n}\n\n- (void)updateTouches:(NSSet*)touches {\n  for (UITouch* touch in touches) {\n    if (touch.view != [self glkView]) {\n      continue;\n    }\n    CGPoint location = [touch locationInView:touch.view];\n    EbitenmobileviewUpdateTouchesOnIOS(touch.phase, (uintptr_t)touch, location.x, location.y);\n  }\n}\n\n- (void)touchesBegan:(NSSet*)touches withEvent:(UIEvent*)event {\n  [self updateTouches:touches];\n}\n\n- (void)touchesMoved:(NSSet*)touches withEvent:(UIEvent*)event {\n  [self updateTouches:touches];\n}\n\n- (void)touchesEnded:(NSSet*)touches withEvent:(UIEvent*)event {\n  [self updateTouches:touches];\n}\n\n- (void)touchesCancelled:(NSSet*)touches withEvent:(UIEvent*)event {\n  [self updateTouches:touches];\n}\n\n@end\n`\n")