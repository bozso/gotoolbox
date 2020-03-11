import utils
import os.path as path
import glob

root = utils.Path(path.dirname(path.abspath(__file__)))

class Project(object):
    flags = "-ldflags '-s -w'"
    
    @staticmethod
    def sources(*args):
        return glob.glob(path.join(*args))

    def generate_ninja(self):
        # src_path = path.join(root, "gamma")
        # src = sources(src_path, "*.go")
        
        subdirs = {path.join(root, elem)
            for elem in {
                "cli", "command", "errors", "path",
            }
        }
        
        
        src = sum(
            (
                self.sources(root, sdir, "*.go")
                for sdir in subdirs
            ),
            []
        )

        main = "toolbox"
        
        
        
        ninja = root.join("build.ninja")
        
        n = utils.Ninja.in_path(root)
        
        cmd = "go build %s -o ${out} ${in}"
        n.rule("go", cmd % self.flags, "Build executable.")
        n.newline()
    
        n.build(
            str(root.join("bin", main)), "go",
            str(root.join(main + ".go")), implicit=src
        )
        
        n.newline()
        
        for sdir in subdirs:
            n = utils.Ninja.in_path(sdir)
            n.subninja(ninja)
    
    
def main():
    p = Project()
    p.generate_ninja()
    
    
    
if __name__ == "__main__":
    main()
