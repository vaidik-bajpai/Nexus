# Nexus

Nexus is a real time collaborative task management application.
My goal with this project is to showcase my ability to build and work with things at scale.
I have written a long and verbose feature list and my goal is to checkmark as many as I can. 

***

# **Nexus - Feature Specification**

---

## **🎯 Core Entities & Data Models**

Before features, let's define what we're working with:

### **1. User**
- ID, email, password (hashed), name, avatar URL
- Role: `admin`, `manager`, `member`
- Status: `online`, `offline`, `away`
- Created/updated timestamps

### **2. Workspace/Team**
- ID, name, description, owner ID
- Members (user IDs with roles)
- Settings (permissions, notifications)

### **3. Project**
- ID, name, description, workspace ID
- Color/icon for visual identification
- Status: `active`, `archived`, `completed`
- Created by, assigned members

### **4. Task**
- ID, title, description, project ID
- Assigned to (user ID), created by (user ID)
- Status: `todo`, `in_progress`, `in_review`, `done`
- Priority: `low`, `medium`, `high`, `urgent`
- Due date, tags/labels
- Attachments (file URLs)
- Created/updated timestamps

### **5. Comment**
- ID, task ID, user ID, content
- Mentions (@username)
- Timestamp

### **6. Activity/Event**
- ID, type (task_created, task_assigned, status_changed, etc.)
- Related entity IDs (task, user, project)
- Metadata (what changed)
- Timestamp

***

## **📋 Phase 1: Foundation (MVP)**
*Goal: Basic CRUD + Authentication - Get it working*

### **Authentication & Authorization**

**User registration with email validation**
- [x] Credentials
- [x] OAuth

**Login with JWT token generation**
- [x] Credentials
- [x] OAuth

- [x] Token refresh mechanism
- [x] Password reset flow (email-based)
- [x] Logout (token invalidation)
- [x] Middleware: Protect routes, verify tokens
- [ ] Role-based access control (RBAC)

### **User Management**
- [ ] Get user profile
- [ ] Update profile (name, avatar)
- [ ] Change password
- [ ] Delete account

### **Workspace Management**
- [x] Create workspace
- [x] Get workspace details
- [ ] Update workspace
- [ ] Delete workspace (owner only)
**List user's workspaces**
- [x] basic
- [ ] paginated

### **Project Management**
- [x] Create project within workspace
- [ ] List projects in workspace
- [ ] Get project details
- [ ] Update project
- [ ] Delete project
- [ ] Archive/unarchive project

### **Task Management (CRUD)**
- [ ] Create task
- [ ] Get task details
- [ ] List tasks (with filters: status, assignee, priority, project)
- [ ] Update task (title, description, status, priority, due date)
- [ ] Delete task
- [ ] Assign/reassign task to user
- [ ] Add/remove tags

### **Basic API Features**
- [ ] Pagination for list endpoints
- [ ] Sorting (by created date, priority, due date)
- [ ] Search tasks by title/description
- [ ] Error handling with proper HTTP status codes
- [ ] Request validation middleware
- [ ] Structured JSON responses

***

## **⚡ Phase 2: Real-Time Features**
*Goal: Add WebSocket support for live updates*

### **WebSocket Server**
- [ ] WebSocket connection handler
- [ ] Authentication via JWT in WebSocket handshake
- [ ] Connection management (store active connections)
- [ ] Heartbeat/ping-pong to detect disconnects
- [ ] Graceful connection cleanup

### **Real-Time Events Broadcasting**
- [ ] **Task Created** → Broadcast to project members
- [ ] **Task Updated** → Broadcast to watchers/assignee
- [ ] **Task Assigned** → Notify assignee instantly
- [ ] **Status Changed** → Broadcast to project members
- [ ] **Comment Added** → Notify task watchers
- [ ] **User Online/Offline** → Broadcast presence to workspace
- [ ] **User Viewing Task** → Show "👁️ Mike is viewing this"

### **Presence System**
- [ ] Track which users are online in workspace
- [ ] Track which task a user is currently viewing
- [ ] Broadcast presence changes (online/offline/viewing)
- [ ] Show active users in workspace

### **Notification System**
- [ ] In-app notification center (unread count)
- [ ] Mark notification as read
- [ ] Clear all notifications
- [ ] Notification types:
  - Task assigned to you
  - Mentioned in comment
  - Task due date approaching
  - Task status changed
  - Comment reply

***

## **🔥 Phase 3: Collaboration Features**
*Goal: Make it truly collaborative*

### **Comments & Mentions**
- [ ] Add comment to task
- [ ] Edit/delete own comments
- [ ] Mention users with `@username`
- [ ] Get comments for task (with pagination)
- [ ] Real-time comment updates

### **File Attachments**
- [ ] Upload files to task (images, PDFs, docs)
- [ ] Store files (local storage or S3/CloudFlare R2)
- [ ] List attachments on task
- [ ] Delete attachment
- [ ] Download attachment

### **Task Dependencies**
- [ ] Mark task as blocked by another task
- [ ] Show dependency chain
- [ ] Warn when trying to complete task with incomplete dependencies

### **Activity Feed**
- [ ] Workspace activity feed (all recent actions)
- [ ] Project activity feed
- [ ] Task activity timeline (who did what, when)
- [ ] Real-time activity updates

### **Team Collaboration**
- [ ] Invite users to workspace (via email)
- [ ] Accept/decline workspace invitations
- [ ] Manage team members (add/remove)
- [ ] Assign roles to members
- [ ] List workspace members with online status

***

## **🚀 Phase 4: Advanced Features**
*Goal: Production-ready, enterprise-level functionality*

### **Performance & Scalability**
- [ ] **Redis Caching**:
  - Cache frequently accessed data (user sessions, workspace details)
  - Cache task lists (invalidate on update)
- [ ] **Rate Limiting**:
  - Per-user rate limits (100 requests/min)
  - Per-IP rate limits for public endpoints
  - WebSocket message rate limiting
- [ ] **Database Optimization**:
  - Indexes on frequently queried fields
  - Query optimization (use EXPLAIN ANALYZE)
  - Connection pooling

### **Background Jobs**
- [ ] Email notifications (task assigned, mentions)
- [ ] Daily digest emails (tasks due today)
- [ ] Cleanup expired tokens
- [ ] Archive old completed tasks
- [ ] Generate weekly reports

### **Analytics & Insights**
- [ ] Task completion rate by user
- [ ] Average task completion time
- [ ] Overdue tasks report
- [ ] User activity metrics
- [ ] Project progress dashboard data
- [ ] Export analytics as CSV/JSON

### **Advanced Task Features**
- [ ] Recurring tasks (daily, weekly, monthly)
- [ ] Task templates
- [ ] Bulk operations (assign multiple, change status for multiple)
- [ ] Custom fields per workspace
- [ ] Task priority auto-adjustment based on due date

### **Search & Filters**
- [ ] Full-text search across tasks, comments
- [ ] Advanced filters:
  - Multiple status selection
  - Date range (created/due)
  - Multiple assignees
  - Tags (AND/OR logic)
- [ ] Save filter presets
- [ ] Search suggestions/autocomplete

***

## **🛡️ Phase 5: Security & DevOps**
*Goal: Make it production-ready and secure*

### **Security**
- [ ] **Input Validation**: Sanitize all user inputs
- [ ] **SQL Injection Prevention**: Use parameterized queries
- [ ] **XSS Prevention**: Escape outputs if serving HTML
- [ ] **CSRF Protection**: For any state-changing operations
- [ ] **Rate Limiting**: Prevent brute force/DoS
- [ ] **Secure Headers**: CORS, CSP, HSTS
- [ ] **Audit Logging**: Log sensitive actions (login, role changes)

### **Testing**
- [ ] Unit tests for business logic (80%+ coverage)
- [ ] Integration tests for API endpoints
- [ ] WebSocket connection tests
- [ ] Load testing (simulate 1000+ concurrent users)
- [ ] Database migration tests

### **DevOps & Deployment**
- [ ] **Docker**: Containerize the application
- [ ] **Docker Compose**: Multi-container setup (app + postgres + redis)
- [ ] **Environment Config**: Use `.env` files, never hardcode secrets
- [ ] **CI/CD Pipeline** (GitHub Actions):
  - Run tests on every push
  - Lint code (golangci-lint)
  - Build Docker image
  - Deploy to staging/production
- [ ] **Monitoring**:
  - Prometheus metrics endpoint
  - Grafana dashboards (request rate, error rate, latency)
  - Logging (structured logs with zerolog or zap)
- [ ] **Database Migrations**: Use `golang-migrate` or similar
- [ ] **Health Check Endpoint**: `/health` for load balancers

### **Documentation**
- [ ] **API Documentation**: Swagger/OpenAPI spec
- [ ] **README**: Setup instructions, architecture overview
- [ ] **Architecture Diagram**: Show components (API, WebSocket, DB, Redis)
- [ ] **Deployment Guide**: How to deploy to production
- [ ] **Contribution Guide**: If open-sourcing

***

## **🎨 Phase 6: Polish & Extras (Optional)**
*Goal: Stand out from other candidates*

### **Nice-to-Have Features**
- [ ] **Dark Mode API Support**: Send theme preferences
- [ ] **Task Time Tracking**: Start/stop timer on tasks
- [ ] **Kanban Board API**: Return tasks grouped by status
- [ ] **Calendar View Data**: Tasks grouped by due date
- [ ] **Webhooks**: Trigger external services on events
- [ ] **API Versioning**: `/api/v1/` for future-proofing
- [ ] **GraphQL API**: Alternative to REST (bonus points)
- [ ] **Mobile Push Notifications**: FCM/APNS integration
- [ ] **Task Export**: Export project tasks as CSV/PDF
- [ ] **Slack Integration**: Post updates to Slack channels

### **Developer Experience**
- [ ] **Postman Collection**: Pre-configured API requests
- [ ] **Sample Data Seeder**: Populate DB with test data
- [ ] **CLI Tool**: Manage tasks from terminal
- [ ] **SDK/Client Library**: Go client for the API

***

## **📊 Feature Priority Matrix**

| **Must Have (Phase 1-2)** | **Should Have (Phase 3)** | **Nice to Have (Phase 4-6)** |
|---------------------------|---------------------------|------------------------------|
| Authentication | Comments & Mentions | Analytics Dashboard |
| Task CRUD | File Attachments | Recurring Tasks |
| WebSocket Real-time | Activity Feed | GraphQL API |
| Presence System | Team Management | Webhooks |
| Basic Notifications | Task Dependencies | Time Tracking |

***

## **🎯 Recommended Build Order**

**Week 1-2:** Phase 1 (Foundation)  
**Week 3:** Phase 2 (WebSocket + Real-time)  
**Week 4:** Phase 3 (Collaboration)  
**Week 5:** Phase 4 (Polish, Performance, Security)  
**Week 6:** Phase 5 (Testing, DevOps, Documentation)  
**Week 7+:** Phase 6 (Extras if time permits)

***

## **Next Steps**

Now that we have the full feature spec:

1. **Would you like me to create a detailed project structure** (folder layout, package organization)?
2. **Should I draft the database schema** with SQL migration files?
3. **Want a Roadmap/Kanban board template** for tracking your own progress?
4. **Need help with the tech stack decisions** (which framework, libraries to use)?

Let me know which part you want to dive into next! 🚀